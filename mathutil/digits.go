package mathutil

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

// Digits returns the characters needed to print the value (the number of
// digits plus potentially a sign marker)
func Digits[T constraints.Signed](v T) int {
	return DigitsInBase(v, 10)
}

// DigitsUnsigned returns the characters needed to print the value
func DigitsUnsigned[T constraints.Unsigned](v T) int {
	return DigitsInBaseUnsigned(v, 10)
}

// DigitsInBase returns the characters needed to print the value v in base
// b. Note that the base must be 2 or more; if not a panic is generated.
func DigitsInBase[T constraints.Signed](v T, b uint) int {
	if b < 2 {
		panic(fmt.Sprintf("Invalid base (%d), the base must be at least 2", b))
	}

	if v == 0 {
		return 1
	}

	var d int
	if v < 0 {
		d++
		v *= -1
	}

	logConv := 1.0
	if b != 10 {
		logConv = math.Log10(float64(b))
	}
	d += int(math.Ceil(math.Log10(float64(v+1)) / logConv))

	return d
}

// DigitsInBaseUnsigned returns the characters needed to print the value v
// (of an unsigned integer type) in base b. Note that the base must be 2 or
// more; if not a panic is generated.
func DigitsInBaseUnsigned[T constraints.Unsigned](v T, b uint) int {
	if b < 2 {
		panic(fmt.Sprintf("Invalid base (%d), the base must be at least 2", b))
	}

	if v == 0 {
		return 1
	}

	logConv := 1.0
	if b != 10 {
		logConv = math.Log10(float64(b))
	}
	return int(math.Ceil(math.Log10(float64(v+1)) / logConv))
}
