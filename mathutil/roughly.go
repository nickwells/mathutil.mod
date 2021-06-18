package mathutil

import (
	"math"
)

// trialRound will round v to the nearest multiple of factor
func trialRound(v, factor float64) float64 {
	v /= factor
	v = math.Round(v)
	return v * factor
}

// Roughly converts v to a value that is "roughly" the same but closer to
// some multiple of five or ten. It will never be more than accuracy percent
// from the original value. The accuracy must be less than 100 and greater
// than zero.
func Roughly(v, accuracy float64) float64 {
	if v == 0 {
		return v
	}
	if accuracy <= 0.0 || accuracy >= 100.0 {
		return v
	}

	accuracy /= 100.0

	signMult := 1.0
	if v < 0 {
		signMult = -1
	}
	newV := v * signMult

	maxDiff := newV * accuracy
	precision := math.Floor(math.Log10(maxDiff)) - 1
	scale := math.Pow(10.0, precision)
	newV /= scale
	maxDiff /= scale

	for _, factor := range []float64{100, 50, 10, 5} {
		trial := trialRound(newV, factor)
		diff := math.Abs(trial - newV)
		if diff < maxDiff {
			newV = trial
			break
		}
	}
	newV *= scale

	return newV * signMult
}
