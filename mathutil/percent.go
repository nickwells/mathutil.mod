package mathutil

import "golang.org/x/exp/constraints"

const percentFactor = 100

// FromPercent takes a percentage value and converts it into a value of the
// same type. So, for instance, a pct value of 4% would be converted into a
// value 0.04 by dividing it by 100.
func FromPercent[T constraints.Float | constraints.Integer](pct T) T {
	return pct / percentFactor
}

// ToPercent takes a value and converts it to a percentage. So, for instance,
// a float of 0.04 would be converted into a percentage value of 4 by
// multiplying it by 100.
func ToPercent[T constraints.Float | constraints.Integer](v T) T {
	return v * percentFactor
}
