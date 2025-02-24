package mathutil

import (
	"errors"
	"fmt"
	"math"
)

var (
	errNumeratorTooBig   = errors.New("overflow: the numerator is too big")
	errDenominatorTooBig = errors.New("overflow: the denominator is too big")
)

// Rational represents a rational number. It is used as the return value of
// the RationalApproximation functions.
type Rational struct {
	N int64
	D int64
}

// String returns a string value for the Rational
func (r Rational) String() string {
	return fmt.Sprintf("Rational{N: %19d, D: %19d}", r.N, r.D)
}

// AsFloat64 returns the float64 equivalent of the Rational
func (r Rational) AsFloat64() float64 {
	return float64(r.N) / float64(r.D)
}

// Invert returns 1/r
func (r Rational) Invert() Rational {
	return Rational{N: r.D, D: r.N}
}

// Proximity returns the absolute difference between the rational value and
// the supplied value as a proportion of the supplied value
func (r Rational) Proximity(v float64) float64 {
	return math.Abs((r.AsFloat64() - v) / v)
}

// mediant generates the Mediant of the two values; given a/b and c/d the
// Mediant is a+c/b+d. If the lower and upper values satisfy the properties
// that a/b < c/d and bc-ad == 1 then the Mediant is guaranteed to lie
// between them. Additionally the new value also satisfies this property with
// respect to lower and upper values respectively. That is, if we take a/b
// and c/d such that a/b < c/d and bc-ad == 1 then p/q where p=a+c and q=b+d
// will also have the property that bp-aq == 1 and qc-pd == 1.
//
// It will return a non-nil error if any of the sums overflow. Note that it
// does not check that the Mediant property holds, this is the responsibility
// of the caller.
func mediant(lower, upper Rational) (Rational, error) {
	if math.MaxInt64-lower.N < upper.N {
		return Rational{0, 1}, errNumeratorTooBig
	}

	if math.MaxInt64-lower.D < upper.D {
		return Rational{0, 1}, errDenominatorTooBig
	}

	return Rational{N: lower.N + upper.N, D: lower.D + upper.D}, nil
}

// SetRational constructs a Rational from the passed values, checks for
// overflows and returns it along with any error
func SetRational(intPart float64, n, d, sign int64) (Rational, error) {
	numerator := float64(d)*intPart + float64(n)
	if numerator > float64(math.MaxInt64) {
		return Rational{}, errTooBig
	}

	return Rational{N: (d*int64(intPart) + n) * sign, D: d}, nil
}
