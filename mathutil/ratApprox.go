package mathutil

import (
	"errors"
	"math"
)

// MaxFareyTrials is the maximum number of attempts that will be made to find
// a Rational Approximation by taking entries in the Farey Sequence. If no
// approximation within the accuracy threshold has been found by this point
// an error is returned indicating that the Farey sequence converges too
// slowly. Conversion can be very slow for values, for instance, with
// fractional parts very close to the ends of the interval [0,1].
const MaxFareyTrials = 100

const raErrSuffix = ", no rational approximation is possible"

var (
	errInaccurate  = errors.New("couldn't meet accuracy target" + raErrSuffix)
	errTooBig      = errors.New("the value is too big" + raErrSuffix)
	errIsInf       = errors.New("the value is infinite" + raErrSuffix)
	errIsNaN       = errors.New("the value is not a number" + raErrSuffix)
	errBadAccuracy = errors.New("accuracy must be >0 and <100" + raErrSuffix)
	errConvTooSlow = errors.New("Farey seq converges too slowly" + raErrSuffix)
)

// RationalApproximationByFareysAlgo returns a Rational approximation for v.
// Various float64 values cannot be represented in this way (including values
// that are too big) and a non-nil error is returned in this case. It uses
// Farey's algorithm to generate a sequence of rational approximations.
//
// The accuracy must be less than 100 and greater than zero.
//
// This may generate different approximations from the RationalApproximation
// func.
//
// Note that this will try at most MaxFareyTrials times before giving up. It
// can be very slow to converge to certain values, particularly those close
// to zero or one. If an error is returned the
func RationalApproximationByFareysAlgo(v, accuracy float64) (Rational, error) {
	var r Rational

	if err := checkRationalApproxParams(v, accuracy); err != nil {
		return r, err
	}

	accuracy = FromPercent(accuracy)

	vAbs, sign := normaliseRationalApproxVal(v)

	intPart := math.Floor(vAbs)
	fracPart := vAbs - intPart

	if fracPart == 0 {
		return Rational{int64(intPart) * sign, 1}, nil
	}

	var (
		lower = Rational{N: 0, D: 1}
		upper = Rational{N: 1, D: 1}
	)

	for range MaxFareyTrials {
		mediant, err := mediant(lower, upper)
		if err != nil {
			return r, errors.New(err.Error() + raErrSuffix)
		}

		r, err = SetRational(intPart, mediant.N, mediant.D, sign)
		if err != nil {
			return r, err
		}

		if r.Proximity(v) <= accuracy {
			return r, nil
		}

		rv := mediant.AsFloat64()
		if fracPart > rv {
			lower = mediant
		} else {
			upper = mediant
		}
	}

	return r, errConvTooSlow
}

// normaliseRationalApproxVal converts v to its absolute value and returns it
// together with a sign value
func normaliseRationalApproxVal(v float64) (float64, int64) {
	vAbs := math.Abs(v)

	var sign int64 = 1

	if math.Signbit(v) {
		sign = -1
	}

	return vAbs, sign
}

// continuedFraction returns a slice containing at most maxVals entries which
// represents the first entries in the continued fractions representation of
// v.
func continuedFraction(v float64, maxVals uint) ([]int64, error) {
	if math.IsInf(v, 1) ||
		math.IsInf(v, -1) {
		return nil, errors.New("the value is infinite")
	}

	if math.IsNaN(v) {
		return nil, errors.New("the value is not a number")
	}

	cf := make([]int64, 0, maxVals)

	for i := uint(0); i < maxVals; i++ {
		intPart := math.Floor(v)
		if intPart > float64(math.MaxInt64) ||
			float64(int64(intPart)) != intPart {
			return cf, errors.New("continued fraction value too big")
		}

		cf = append(cf, int64(intPart))

		v -= intPart

		if v <= 0 {
			break
		}

		v = 1 / v
	}

	return cf, nil
}

// checkRationalAccuracy returns a non nil error if accuracy <= 0 or >= 100
func checkRationalAccuracy(accuracy float64) error {
	if accuracy <= 0.0 || accuracy >= 100.0 {
		return errBadAccuracy
	}

	return nil
}

// checkRationalTargetVal returns a non-nil error if the value is not a
// regular float64 value
func checkRationalTargetVal(v float64) error {
	if math.IsInf(v, 1) ||
		math.IsInf(v, -1) {
		return errIsInf
	}

	if math.IsNaN(v) {
		return errIsNaN
	}

	if v >= float64(math.MaxInt64) ||
		v < float64(math.MinInt64) {
		return errTooBig
	}

	return nil
}

// checkRationalApproxParams checks that the parameters to
// RationalApproximation are valid and returns an error if not
func checkRationalApproxParams(v float64, accuracy float64) error {
	if err := checkRationalAccuracy(accuracy); err != nil {
		return err
	}

	if err := checkRationalTargetVal(v); err != nil {
		return err
	}

	return nil
}

// RationalApproximation returns a Rational value (a numerator N and
// denominator D) and an error. The values are such that the supplied value v
// will lie within accuracy percent of N/D. It uses continued fractions to
// generate the rational coefficients. Various float64 values cannot be
// represented in this way (including values that are too big) and a non-nil
// error is returned in this case.
//
// The accuracy is expressed as a percentage and must be less than 100 and
// greater than zero.
//
// This may generate different approximations from the
// RationalApproximationByFareysAlgo func.
func RationalApproximation(v, accuracy float64) (Rational, error) {
	if v == 0 {
		return Rational{N: 0, D: 1}, nil
	}

	var r Rational

	if err := checkRationalApproxParams(v, accuracy); err != nil {
		return r, err
	}

	accuracy = FromPercent(accuracy)

	vAbs, sign := normaliseRationalApproxVal(v)

	const maxCFLen = 20

	for cfLen := uint(1); cfLen <= maxCFLen; cfLen++ {
		cf, err := continuedFraction(vAbs, cfLen)
		if err != nil && len(cf) == 0 {
			return r, errors.New(err.Error() + raErrSuffix)
		}

		idxEnd := len(cf) - 1
		r = Rational{N: 1, D: cf[idxEnd]}
		idxEnd--

		for ; idxEnd >= 0; idxEnd-- {
			if math.MaxInt64/r.D < cf[idxEnd] { // restart, ignoring end values
				r = Rational{N: 1, D: cf[idxEnd]}
				continue
			}

			p := cf[idxEnd] * r.D

			if math.MaxInt64-p < r.N { // restart, ignoring end values
				r = Rational{N: 1, D: cf[idxEnd]}
				continue
			}

			r.N += p
			r = r.Invert()
		}

		r = r.Invert() // undo the last swap

		r.N *= sign
		if r.Proximity(v) <= accuracy {
			return r, nil
		}
	}

	return r, errInaccurate
}
