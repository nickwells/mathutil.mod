package mathutil

import "math"

// WithinNPercent returns true if a and b are within epsilon percent of one
// another.  Strictly speaking the test is for whether the difference between
// a and b as a proportion of the larger value is less than epsilon. So an
// epsilon of 10 will test for numbers within 10% of each other.  Numbers
// with differing sign are always considered different regardless of
// proximity. The epsilon value is forced to a positive value (the absolute
// value is taken).
func WithinNPercent(a, b, epsilon float64) bool {
	if a == b {
		return true
	}

	absA := math.Abs(a)
	absB := math.Abs(b)

	isPosA := (a == absA)
	isPosB := (b == absB)

	if isPosA != isPosB {
		return false
	}

	pctDiff := ToPercent(math.Abs(absA-absB) / math.Max(absA, absB))

	return pctDiff <= math.Abs(epsilon)
}

// AlmostEqual returns true if a and b are within epsilon of one another
// regardless of sign. The epsilon value is forced to a positive value (the
// absolute value is taken).
func AlmostEqual(a, b, epsilon float64) bool {
	if a == b {
		return true
	}

	return math.Abs(a-b) < math.Abs(epsilon)
}
