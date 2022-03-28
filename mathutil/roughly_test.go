package mathutil_test

import (
	"testing"

	"github.com/nickwells/mathutil.mod/v2/mathutil"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestRoughly(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		v       float64
		pct     float64
		expVal  float64
		epsilon float64
	}{
		{
			ID:      testhelper.MkID("< 1% - round to unit"),
			v:       123.456,
			pct:     1.0,
			expVal:  123.0,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("< 2% - round to x5"),
			v:       123.456,
			pct:     2.0,
			expVal:  125.0,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("< 3% - round to x10"),
			v:       123.456,
			pct:     3.0,
			expVal:  120.0,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("negative, < 1% - round to unit"),
			v:       -123.456,
			pct:     1.0,
			expVal:  -123.0,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("negative, < 2% - round to x5"),
			v:       -123.456,
			pct:     2.0,
			expVal:  -125.0,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("negative, < 3% - round to x10"),
			v:       -123.456,
			pct:     3.0,
			expVal:  -120.0,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("small, < 1% - round to unit"),
			v:       0.00123456,
			pct:     1.0,
			expVal:  0.00123,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("small, < 2% - round to x5"),
			v:       0.00123456,
			pct:     2.0,
			expVal:  0.00125,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("small, < 3% - round to x10"),
			v:       0.00123456,
			pct:     3.0,
			expVal:  0.00120,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("small, negative, < 1% - round to unit"),
			v:       -0.00123456,
			pct:     1.0,
			expVal:  -0.00123,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("small, negative, < 2% - round to x5"),
			v:       -0.00123456,
			pct:     2.0,
			expVal:  -0.00125,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("small, negative, < 3% - round to x10"),
			v:       -0.00123456,
			pct:     3.0,
			expVal:  -0.00120,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("large, < 1% - round to unit"),
			v:       12345600.0,
			pct:     1.0,
			expVal:  12300000,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("large, < 2% - round to x5"),
			v:       12345600.0,
			pct:     2.0,
			expVal:  12500000,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("large, < 3% - round to x10"),
			v:       12345600.0,
			pct:     3.0,
			expVal:  12000000,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("large, negative, < 1% - round to unit"),
			v:       -12345600.0,
			pct:     1.0,
			expVal:  -12300000,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("large, negative, < 2% - round to x5"),
			v:       -12345600.0,
			pct:     2.0,
			expVal:  -12500000,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("large, negative, < 3% - round to x10"),
			v:       -12345600.0,
			pct:     3.0,
			expVal:  -12000000,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("zero"),
			v:       0.0,
			pct:     1.0,
			expVal:  0.0,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("bad %age - negative"),
			v:       123.456,
			pct:     -1.0,
			expVal:  123.456,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("bad %age - too big"),
			v:       123.456,
			pct:     100.0,
			expVal:  123.456,
			epsilon: 0.0000001,
		},
		{
			ID:      testhelper.MkID("bad %age - zero"),
			v:       123.456,
			pct:     0.0,
			expVal:  123.456,
			epsilon: 0.0000001,
		},
	}

	for _, tc := range testCases {
		r := mathutil.Roughly(tc.v, tc.pct)
		testhelper.DiffFloat(t, tc.IDStr(), "", r, tc.expVal, tc.epsilon)
	}
}
