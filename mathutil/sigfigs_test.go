package mathutil_test

import (
	"testing"

	"github.com/nickwells/mathutil.mod/v2/mathutil"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestDigitsForSigFigs(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpPanic
		sigFigs  uint8
		v        float64
		expWidth int
		expPrec  int
	}{
		{
			ID: testhelper.MkID("bad sig figs"),
			ExpPanic: testhelper.MkExpPanic(
				"the number of significant figures must be greater than zero"),
			sigFigs: 0,
		},
		{
			ID:       testhelper.MkID("zero, 1SF"),
			sigFigs:  1,
			v:        0,
			expWidth: 1,
			expPrec:  0,
		},
		{
			ID:       testhelper.MkID("zero, 2SF"),
			sigFigs:  2,
			v:        0,
			expWidth: 3,
			expPrec:  1,
		},
		{
			ID:       testhelper.MkID("one, 1SF"),
			sigFigs:  1,
			v:        1,
			expWidth: 1,
			expPrec:  0,
		},
		{
			ID:       testhelper.MkID("one, 2SF"),
			sigFigs:  2,
			v:        1,
			expWidth: 3,
			expPrec:  1,
		},
		{
			ID:       testhelper.MkID("99, 1SF"),
			sigFigs:  1,
			v:        99,
			expWidth: 2,
			expPrec:  0,
		},
		{
			ID:       testhelper.MkID("99, 2SF"),
			sigFigs:  2,
			v:        99,
			expWidth: 2,
			expPrec:  0,
		},
		{
			ID:       testhelper.MkID("99, 3SF"),
			sigFigs:  3,
			v:        99,
			expWidth: 4,
			expPrec:  1,
		},
		{
			ID:       testhelper.MkID("0.5, 1SF"),
			sigFigs:  1,
			v:        0.5,
			expWidth: 3,
			expPrec:  1,
		},
		{
			ID:       testhelper.MkID("0.5, 2SF"),
			sigFigs:  2,
			v:        0.5,
			expWidth: 4,
			expPrec:  2,
		},
		{
			ID:       testhelper.MkID("0.05, 1SF"),
			sigFigs:  1,
			v:        0.05,
			expWidth: 4,
			expPrec:  2,
		},
		{
			ID:       testhelper.MkID("0.05, 2SF"),
			sigFigs:  2,
			v:        0.05,
			expWidth: 5,
			expPrec:  3,
		},
	}

	for _, tc := range testCases {
		panicked, panicVal := testhelper.PanicSafe(func() {
			w, p := mathutil.FmtValsForSigFigs(tc.sigFigs, tc.v)
			testhelper.DiffInt(t, tc.IDStr(), "width", w, tc.expWidth)
			testhelper.DiffInt(t, tc.IDStr(), "precision", p, tc.expPrec)

			if tc.v == 0.0 {
				return
			}
			w, p = mathutil.FmtValsForSigFigs(tc.sigFigs, tc.v*-1.0)
			testhelper.DiffInt(t,
				tc.IDStr()+" (-ve value)", "width", w, tc.expWidth+1)
			testhelper.DiffInt(t,
				tc.IDStr()+" (-ve value)", "precision", p, tc.expPrec)
		})
		testhelper.CheckExpPanic(t, panicked, panicVal, tc)
	}
}

func TestDigitsForSigFigsMulti(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpPanic
		sigFigs  uint8
		v        float64
		vals     []float64
		expWidth int
		expPrec  int
	}{
		{
			ID: testhelper.MkID("bad sig figs"),
			ExpPanic: testhelper.MkExpPanic(
				"the number of significant figures must be greater than zero"),
			sigFigs: 0,
		},
		{
			ID:       testhelper.MkID("one val, SF: 1"),
			sigFigs:  1,
			v:        1,
			expWidth: 1,
			expPrec:  0,
		},
		{
			ID:       testhelper.MkID("2 vals, > and < 1, SF: 1"),
			sigFigs:  1,
			v:        10,
			vals:     []float64{0.5},
			expWidth: 4,
			expPrec:  1,
		},
		{
			ID:       testhelper.MkID("2 vals, > and < 1, -ve, SF: 1"),
			sigFigs:  1,
			v:        1,
			vals:     []float64{-0.5},
			expWidth: 4,
			expPrec:  1,
		},
		{
			ID:       testhelper.MkID("many vals, SF: 2"),
			sigFigs:  1,
			v:        1,
			vals:     []float64{20, 0.001, -0.5},
			expWidth: 7,
			expPrec:  4,
		},
	}

	for _, tc := range testCases {
		panicked, panicVal := testhelper.PanicSafe(func() {
			w, p := mathutil.FmtValsForSigFigsMulti(tc.sigFigs, tc.v, tc.vals...)
			testhelper.DiffInt(t, tc.IDStr(), "width", w, tc.expWidth)
			testhelper.DiffInt(t, tc.IDStr(), "precision", p, tc.expPrec)
		})
		testhelper.CheckExpPanic(t, panicked, panicVal, tc)
	}
}
