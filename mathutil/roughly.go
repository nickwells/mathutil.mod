package mathutil

import (
	"math"

	"golang.org/x/exp/constraints"
)

// trialRound will round v to the nearest multiple of factor
func trialRound[F constraints.Float](v F, factor float64) F {
	rounded := float64(v)
	rounded /= factor
	rounded = math.Round(rounded)

	return F(rounded * factor)
}

// Roughly converts v to a value that is "roughly" the same but closer to
// some multiple of five or ten. It will never be more than accuracy percent
// from the original value. The accuracy must be less than 100 and greater
// than zero.
func Roughly[F constraints.Float](v, accuracy F) F {
	if v == 0 {
		return v
	}

	if accuracy <= 0.0 || accuracy >= 100.0 {
		return v
	}

	accuracy = FromPercent(accuracy)

	var signMult F = 1.0
	if v < 0 {
		signMult = -1
	}

	newV := v * F(signMult)

	maxDiff := newV * accuracy
	precision := math.Floor(math.Log10(float64(maxDiff))) - 1
	scale := math.Pow(base10, precision)
	newV /= F(scale)
	maxDiff /= F(scale)

	for _, factor := range []float64{100, 50, 10, 5} {
		trial := trialRound(newV, factor)
		diff := math.Abs(float64(trial - newV))

		if F(diff) < maxDiff {
			newV = trial
			break
		}
	}

	newV *= F(scale)

	return newV * signMult
}
