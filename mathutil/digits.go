package mathutil

import (
	"fmt"
	"math"
)

// Digits returns the characters needed to print the value (the number of
// digits plus potentially a sign marker)
func Digits(v int64) int {
	if v == 0 {
		return 1
	}

	d := 0
	if v < 0 {
		d++
		v *= -1
	}
	d += int(math.Ceil(math.Log10(float64(v + 1))))

	return d
}

// DigitsInBase returns the characters needed to print the value v in base
// b. Note that the base must be 2 or more; if not a panic is generated.
func DigitsInBase(v int64, b uint) int {
	if b < 2 {
		panic(fmt.Sprintf("Invalid base (%d), the base must be at least 2", b))
	}

	if v == 0 {
		return 1
	}

	d := 0
	if v < 0 {
		d++
		v *= -1
	}
	d += int(math.Ceil(math.Log10(float64(v+1)) / math.Log10(float64(b))))

	return d
}
