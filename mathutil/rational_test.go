package mathutil

import (
	"math"
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestMediant(t *testing.T) {
	const bigVal = 3 * math.MaxInt64 / 4
	testCases := []struct {
		testhelper.ID
		lower  Rational
		upper  Rational
		expMed Rational
		testhelper.ExpErr
	}{
		{
			ID:     testhelper.MkID("0-1, 1/2"),
			lower:  Rational{0, 1},
			upper:  Rational{1, 1},
			expMed: Rational{1, 2},
		},
		{
			ID:     testhelper.MkID("numerator overflow"),
			lower:  Rational{bigVal, bigVal + 1},
			upper:  Rational{bigVal + 1, bigVal + 2},
			expMed: Rational{0, 1},
			ExpErr: testhelper.MkExpErr(errNumeratorTooBig.Error()),
		},
		{
			ID:     testhelper.MkID("denominator overflow"),
			lower:  Rational{1, bigVal + 2},
			upper:  Rational{1, bigVal + 1},
			expMed: Rational{0, 1},
			ExpErr: testhelper.MkExpErr(errDenominatorTooBig.Error()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.IDStr(), func(t *testing.T) {
			m, err := mediant(tc.lower, tc.upper)
			testhelper.DiffInt(t, tc.IDStr(), "Numerator", m.N, tc.expMed.N)
			testhelper.DiffInt(t, tc.IDStr(), "Denominator", m.D, tc.expMed.D)
			testhelper.CheckExpErr(t, err, tc)
		})
	}
}
