package mathutil

import (
	"math"
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestRationalApproximation(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		v        float64
		accuracy float64
		expRat   Rational
		testhelper.ExpErr
	}{
		{
			ID:       testhelper.MkID("zero"),
			v:        0,
			accuracy: 1,
			expRat:   Rational{N: 0, D: 1},
		},
		{
			ID:       testhelper.MkID("one"),
			v:        1,
			accuracy: 1,
			expRat:   Rational{N: 1, D: 1},
		},
		{
			ID:       testhelper.MkID("minus one"),
			v:        -1,
			accuracy: 1,
			expRat:   Rational{N: -1, D: 1},
		},
		{
			ID:       testhelper.MkID("0.65 to 10%"),
			v:        0.65,
			accuracy: 1e1,
			expRat:   Rational{N: 2, D: 3},
		},
		{
			ID:       testhelper.MkID("0.65 to 0.1%"),
			v:        0.65,
			accuracy: 1e-1,
			expRat:   Rational{N: 13, D: 20},
		},
		{
			ID:       testhelper.MkID("0.65, very accurate"),
			v:        0.65,
			accuracy: 1e-20,
			expRat:   Rational{N: 13, D: 20},
		},
		{
			ID:       testhelper.MkID("small value, very accurate"),
			v:        1.23e-10,
			accuracy: 1e-20,
			expRat:   Rational{N: 123, D: 1000000000000},
		},
		{
			ID:       testhelper.MkID("very small value, very accurate"),
			v:        1.23e-20,
			accuracy: 1e-20,
			expRat:   Rational{N: 0, D: 1},
			ExpErr:   testhelper.MkExpErr(errInaccurate.Error()),
		},
		{
			ID:       testhelper.MkID("Pi to 1%"),
			v:        math.Pi,
			accuracy: 1,
			expRat:   Rational{N: 22, D: 7},
		},
		{
			ID:       testhelper.MkID("Pi to 0.001%"),
			v:        math.Pi,
			accuracy: 0.001,
			expRat:   Rational{N: 355, D: 113},
		},
		{
			ID:       testhelper.MkID("Pi to maximum accuracy"),
			v:        math.Pi,
			accuracy: math.SmallestNonzeroFloat64,
			expRat:   Rational{N: 245850922, D: 78256779},
		},
		{
			ID:       testhelper.MkID("MaxInt64"),
			v:        math.MaxInt64,
			accuracy: 1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errTooBig.Error()),
		},
		{
			ID:       testhelper.MkID("+ve infinity"),
			v:        math.Inf(1),
			accuracy: 1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errIsInf.Error()),
		},
		{
			ID:       testhelper.MkID("-ve infinity"),
			v:        math.Inf(-1),
			accuracy: 1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errIsInf.Error()),
		},
		{
			ID:       testhelper.MkID("NaN"),
			v:        math.NaN(),
			accuracy: 1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errIsNaN.Error()),
		},
		{
			ID:       testhelper.MkID("bad accuracy - negative"),
			v:        1,
			accuracy: -1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errBadAccuracy.Error()),
		},
		{
			ID:       testhelper.MkID("bad accuracy - zero"),
			v:        1,
			accuracy: 0,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errBadAccuracy.Error()),
		},
		{
			ID:       testhelper.MkID("bad accuracy - 100"),
			v:        1,
			accuracy: 100,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errBadAccuracy.Error()),
		},
		{
			ID:       testhelper.MkID("bad accuracy - too big"),
			v:        1,
			accuracy: 101,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errBadAccuracy.Error()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.IDStr(), func(t *testing.T) {
			r, err := RationalApproximation(tc.v, tc.accuracy)
			testhelper.DiffInt(t, tc.IDStr(), "Numerator", r.N, tc.expRat.N)
			testhelper.DiffInt(t, tc.IDStr(), "Denominator", r.D, tc.expRat.D)
			testhelper.CheckExpErr(t, err, tc)
		})
	}
}

func TestRationalApproximationByFareysAlgo(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		v        float64
		accuracy float64
		expRat   Rational
		testhelper.ExpErr
	}{
		{
			ID:       testhelper.MkID("zero"),
			v:        0,
			accuracy: 1,
			expRat:   Rational{N: 0, D: 1},
		},
		{
			ID:       testhelper.MkID("one"),
			v:        1,
			accuracy: 1,
			expRat:   Rational{N: 1, D: 1},
		},
		{
			ID:       testhelper.MkID("minus one"),
			v:        -1,
			accuracy: 1,
			expRat:   Rational{N: -1, D: 1},
		},
		{
			ID:       testhelper.MkID("0.65 to 10%"),
			v:        0.65,
			accuracy: 1e1,
			expRat:   Rational{N: 2, D: 3},
		},
		{
			ID:       testhelper.MkID("0.65 to 0.1%"),
			v:        0.65,
			accuracy: 1e-1,
			expRat:   Rational{N: 13, D: 20},
		},
		{
			ID:       testhelper.MkID("0.65, very accurate"),
			v:        0.65,
			accuracy: 1e-20,
			expRat:   Rational{N: 13, D: 20},
		},
		{
			ID:       testhelper.MkID("small value, very accurate"),
			v:        1.23e-10,
			accuracy: 1e-20,
			expRat:   Rational{N: 1, D: 101},
			ExpErr:   testhelper.MkExpErr(errConvTooSlow.Error()),
		},
		{
			ID:       testhelper.MkID("very small value, very accurate"),
			v:        1.23e-20,
			accuracy: 1e-20,
			expRat:   Rational{N: 1, D: 101},
			ExpErr:   testhelper.MkExpErr(errConvTooSlow.Error()),
		},
		{
			ID:       testhelper.MkID("Pi to 1%"),
			v:        math.Pi,
			accuracy: 1,
			expRat:   Rational{N: 19, D: 6},
		},
		{
			ID:       testhelper.MkID("Pi to 0.001%"),
			v:        math.Pi,
			accuracy: 0.001,
			expRat:   Rational{N: 355, D: 113},
		},
		{
			ID:       testhelper.MkID("Pi to maximum accuracy"),
			v:        math.Pi,
			accuracy: math.SmallestNonzeroFloat64,
			expRat:   Rational{N: 28023, D: 8920},
			ExpErr:   testhelper.MkExpErr(errConvTooSlow.Error()),
		},
		{
			ID:       testhelper.MkID("MaxInt64"),
			v:        math.MaxInt64,
			accuracy: 1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errTooBig.Error()),
		},
		{
			ID:       testhelper.MkID("+ve infinity"),
			v:        math.Inf(1),
			accuracy: 1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errIsInf.Error()),
		},
		{
			ID:       testhelper.MkID("-ve infinity"),
			v:        math.Inf(-1),
			accuracy: 1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errIsInf.Error()),
		},
		{
			ID:       testhelper.MkID("NaN"),
			v:        math.NaN(),
			accuracy: 1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errIsNaN.Error()),
		},
		{
			ID:       testhelper.MkID("bad accuracy - negative"),
			v:        1,
			accuracy: -1,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errBadAccuracy.Error()),
		},
		{
			ID:       testhelper.MkID("bad accuracy - zero"),
			v:        1,
			accuracy: 0,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errBadAccuracy.Error()),
		},
		{
			ID:       testhelper.MkID("bad accuracy - 100"),
			v:        1,
			accuracy: 100,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errBadAccuracy.Error()),
		},
		{
			ID:       testhelper.MkID("bad accuracy - too big"),
			v:        1,
			accuracy: 101,
			expRat:   Rational{N: 0, D: 0},
			ExpErr:   testhelper.MkExpErr(errBadAccuracy.Error()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.IDStr(), func(t *testing.T) {
			r, err := RationalApproximationByFareysAlgo(tc.v, tc.accuracy)
			testhelper.DiffInt(t, tc.IDStr(), "Numerator", r.N, tc.expRat.N)
			testhelper.DiffInt(t, tc.IDStr(), "Denominator", r.D, tc.expRat.D)
			testhelper.CheckExpErr(t, err, tc)
		})
	}
}

func TestContinuedFraction(t *testing.T) {
	const errText = "continued fraction value too big"

	testCases := []struct {
		testhelper.ID
		v       float64
		maxVals uint
		expVals []int64
		testhelper.ExpErr
	}{
		{
			ID:      testhelper.MkID("Pi, 5 vals"),
			v:       math.Pi,
			maxVals: 5,
			expVals: []int64{3, 7, 15, 1, 292},
		},
		{
			ID:      testhelper.MkID("0.65, 13 vals"),
			v:       0.65,
			maxVals: 13,
			expVals: []int64{
				0, 1, 1, 1, 6,
				46912496118442, 1, 1, 1, 42,
				6701785159777, 1, 1,
			},
		},
		{
			ID:      testhelper.MkID("big val, 2 vals"),
			v:       2 * math.MaxInt64,
			maxVals: 2,
			expVals: []int64{},
			ExpErr:  testhelper.MkExpErr(errText),
		},
		{
			ID:      testhelper.MkID("small fraction>1, 2 vals"),
			v:       5 + 1e-18,
			maxVals: 2,
			expVals: []int64{5},
		},
		{
			ID:      testhelper.MkID("very small fraction>1, 2 vals"),
			v:       5 + 1e-19,
			maxVals: 2,
			expVals: []int64{5},
		},
		{
			ID:      testhelper.MkID("very small fraction<1, 2 vals"),
			v:       1e-19,
			maxVals: 2,
			expVals: []int64{0},
			ExpErr:  testhelper.MkExpErr(errText),
		},
		{
			ID:      testhelper.MkID("very small fraction, 200 vals"),
			v:       1e-100,
			maxVals: 200,
			expVals: []int64{0},
			ExpErr:  testhelper.MkExpErr(errText),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.IDStr(), func(t *testing.T) {
			vals, err := continuedFraction(tc.v, tc.maxVals)
			testhelper.DiffSlice(t, tc.IDStr(), "vals", vals, tc.expVals)
			testhelper.CheckExpErr(t, err, tc)
		})
	}
}
