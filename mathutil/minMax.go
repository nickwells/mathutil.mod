package mathutil

import (
	"math"

	"golang.org/x/exp/constraints"
)

// MinOf returns the lesser of the slice of values; it will panic if the
// slice is empty
func MinOf[F constraints.Float](vals ...F) F {
	m := float64(vals[0])
	for _, v := range vals[1:] {
		m = math.Min(m, float64(v))
	}

	return F(m)
}

// MaxOf returns the greater of the slice of values; it will panic if the
// slice is empty
func MaxOf[F constraints.Float](vals ...F) F {
	m := float64(vals[0])
	for _, v := range vals[1:] {
		m = math.Max(m, float64(v))
	}

	return F(m)
}

// MinOfInt returns the lesser of the slice of values; it will panic if the
// slice is empty
func MinOfInt[I constraints.Integer](vals ...I) I {
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}

	return m
}

// MaxOfInt returns the greater of the slice of values; it will panic if the
// slice is empty
func MaxOfInt[I constraints.Integer](vals ...I) I {
	m := vals[0]
	for _, v := range vals[1:] {
		if v > m {
			m = v
		}
	}

	return m
}
