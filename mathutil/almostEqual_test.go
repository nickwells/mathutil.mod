package mathutil_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/nickwells/mathutil.mod/v2/mathutil"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

// TestAlmostEqual tests the AlmostEqual function
func TestAlmostEqual(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		a, b, epsilon float64
		expResult     bool
	}{
		{
			ID:        testhelper.MkID("both positive, true"),
			a:         1.2345,
			b:         1.23456,
			epsilon:   0.0001,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("both positive, false"),
			a:         1.2345,
			b:         1.23456,
			epsilon:   0.00001,
			expResult: false,
		},
		{
			ID:        testhelper.MkID("both negative, true"),
			a:         -1.2345,
			b:         -1.23456,
			epsilon:   0.0001,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("both negative, false"),
			a:         -1.2345,
			b:         -1.23456,
			epsilon:   0.00001,
			expResult: false,
		},
		{
			ID:        testhelper.MkID("different sign, true"),
			a:         -0.00000001,
			b:         0.00000001,
			epsilon:   0.00001,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("identical"),
			a:         1.23456789,
			b:         1.23456789,
			epsilon:   0.00000001,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("max positive"),
			a:         math.MaxFloat64,
			b:         math.MaxFloat64 - 1,
			epsilon:   0.00000000000001,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("max negative"),
			a:         -math.MaxFloat64,
			b:         -math.MaxFloat64 + 1,
			epsilon:   0.00000000000001,
			expResult: true,
		},
	}

	for _, tc := range testCases {
		res := mathutil.AlmostEqual(tc.a, tc.b, tc.epsilon)
		testhelper.DiffBool(t,
			tc.IDStr(),
			fmt.Sprintf("result of AlmostEqual(%.7f, %.7f, %.9f)",
				tc.a, tc.b, tc.epsilon),
			res, tc.expResult)
	}
}

func TestWithinNPercent(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		a, b, pct float64
		expResult bool
	}{
		{
			ID:        testhelper.MkID("both positive, true"),
			a:         100.1,
			b:         100.0,
			pct:       0.15,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("both positive, false"),
			a:         100.3,
			b:         100.0,
			pct:       0.15,
			expResult: false,
		},
		{
			ID:        testhelper.MkID("both negative, true"),
			a:         -100.1,
			b:         -100.0,
			pct:       0.15,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("both negative, false"),
			a:         -100.3,
			b:         -100.0,
			pct:       0.15,
			expResult: false,
		},
		{
			ID:        testhelper.MkID("identical positive"),
			a:         100.0,
			b:         100.0,
			pct:       0.0,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("identical negative"),
			a:         -100.0,
			b:         -100.0,
			pct:       0.0,
			expResult: true,
		},
		{
			ID:        testhelper.MkID("identical magnitude, opposite signs"),
			a:         100.0,
			b:         -100.0,
			pct:       1.0,
			expResult: false,
		},
		{
			ID:        testhelper.MkID("difference equals target %age"),
			a:         100.0,
			b:         99.9,
			pct:       0.1,
			expResult: true,
		},
	}

	for _, tc := range testCases {
		res := mathutil.WithinNPercent(tc.a, tc.b, tc.pct)
		testhelper.DiffBool(t,
			tc.IDStr(),
			fmt.Sprintf("result of WithinNPercent(%11.7f, %11.7f, %5.2f%%)",
				tc.a, tc.b, tc.pct),
			res, tc.expResult)
	}
}
