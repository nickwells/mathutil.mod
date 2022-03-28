package mathutil

import (
	"math"

	"golang.org/x/exp/constraints"
)

// MinOf returns the lesser of the slice of values; it will panic if the
// slice is empty
func MinOf[F constraints.Float](vals ...F) F {
	min := float64(vals[0])
	for _, v := range vals[1:] {
		min = math.Min(min, float64(v))
	}
	return F(min)
}

// MaxOf returns the greater of the slice of values; it will panic if the
// slice is empty
func MaxOf[F constraints.Float](vals ...F) F {
	max := float64(vals[0])
	for _, v := range vals[1:] {
		max = math.Max(max, float64(v))
	}
	return F(max)
}

// MinOfInt returns the lesser of the slice of values; it will panic if the
// slice is empty
func MinOfInt[I constraints.Integer](vals ...I) I {
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// MaxOfInt returns the greater of the slice of values; it will panic if the
// slice is empty
func MaxOfInt[I constraints.Integer](vals ...I) I {
	max := vals[0]
	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
