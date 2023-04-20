package mathutil

import "golang.org/x/exp/constraints"

// FmtValsForSigFigs returns the width and precision needed to print v to at
// least sf significant figures. Note that sf must be greater than 0, a panic
// is generated if not.
//
// There are several caveats to the use of this function: 1, it will only
// generate precision down to 9 digits; 2, for values very close to negative
// powers of 10 the results may appear incorrect and either too many or too
// few digits will be shown. This last caveat is as a result of the way
// floating point numbers are represented by computers and how Go handles
// constants. For instance 0.1*0.1 is not equal to 0.01.
func FmtValsForSigFigs[T constraints.Float](sf uint8, v T) (
	width, precision int,
) {
	if sf == 0 {
		panic("the number of significant figures must be greater than zero")
	}

	if v == 0.0 {
		width = 1
		precision = 0
		if sf > 1 {
			precision = int(sf) - 1
			width++ // for the "."
			width += precision
		}
		return
	}

	digitsPreDP := 1 // always at least 1 digit before the decimal point
	extraWidth := 0
	if v < 0 {
		extraWidth = 1 // for the minus sign
		v *= -1
	}

	if v < 1 {
		const minVal = 1e-9

		extraWidth++ // for the "."
		precision = 1
		for p := T(0.1); v < p && p > minVal; p *= 0.1 {
			precision++
		}
		precision += int(sf) - 1
	} else {
		for p := T(10.0); v >= p; p *= 10 {
			digitsPreDP++
		}
		if digitsPreDP < int(sf) {
			extraWidth++ // for the "."
			precision = int(sf) - digitsPreDP
		}
	}
	width = digitsPreDP + precision + extraWidth
	return
}

// FmtValsForSigFigsMulti returns the width and precision suitable to display
// all the values to at least sf significant figures. For instance 3
// significant figures for the pair 100.0 and 0.1 would require 3 digits
// before the point and 3 digits after the point, with a width of 7 and a
// precision of 3. Note that sf must be greater than 0, a panic is generated
// if not.
func FmtValsForSigFigsMulti[T constraints.Float](sf uint8, v T, vals ...T) (
	width, precision int,
) {
	wid, prec := FmtValsForSigFigs(sf, v)
	precision = prec

	width = digitsBeforePoint(wid, prec)
	for _, val := range vals {
		wid, prec = FmtValsForSigFigs(sf, val)
		if prec > precision {
			precision = prec
		}

		dbp := digitsBeforePoint(wid, prec)
		if dbp > width {
			width = dbp
		}
	}

	width += precision
	if precision > 0 {
		width++
	}
	return
}

// digitsBeforePoint returns the implied number of digits before the decimal
// point implied by the given width and precision.
func digitsBeforePoint(width, precision int) int {
	before := width - precision
	if precision > 0 {
		before-- // for the "."
	}
	return before
}
