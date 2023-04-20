package mathutil_test

import (
	"testing"

	"github.com/nickwells/mathutil.mod/v2/mathutil"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestDigits(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		v         int64
		expDigits int
	}{
		{
			ID:        testhelper.MkID("zero"),
			v:         0,
			expDigits: 1,
		},
		{
			ID:        testhelper.MkID("small"),
			v:         1,
			expDigits: 1,
		},
		{
			ID:        testhelper.MkID("large"),
			v:         9999999,
			expDigits: 7,
		},
	}

	for _, tc := range testCases {
		digits := mathutil.Digits(tc.v)
		testhelper.DiffInt(
			t, tc.IDStr(), "digits", digits, tc.expDigits)
		if tc.v == 0 {
			continue
		}
		digits = mathutil.Digits(tc.v * -1)
		testhelper.DiffInt(
			t, tc.IDStr()+" -ve", "digits", digits, tc.expDigits+1)
	}
}

func TestDigitsUnsigned(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		v         uint64
		expDigits int
	}{
		{
			ID:        testhelper.MkID("zero"),
			v:         0,
			expDigits: 1,
		},
		{
			ID:        testhelper.MkID("small"),
			v:         1,
			expDigits: 1,
		},
		{
			ID:        testhelper.MkID("large"),
			v:         9999999,
			expDigits: 7,
		},
	}

	for _, tc := range testCases {
		digits := mathutil.DigitsUnsigned(tc.v)
		testhelper.DiffInt(t, tc.IDStr(), "digits", digits, tc.expDigits)
	}
}

func TestDigitsInBase(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpPanic
		v         int64
		base      uint
		expDigits int
	}{
		{
			ID: testhelper.MkID("bad base (0)"),
			ExpPanic: testhelper.MkExpPanic(
				"Invalid base (0), the base must be at least 2"),
			base: 0,
		},
		{
			ID: testhelper.MkID("bad base (1)"),
			ExpPanic: testhelper.MkExpPanic(
				"Invalid base (1), the base must be at least 2"),
			base: 1,
		},
		{
			ID:        testhelper.MkID("v: 0, base 2"),
			v:         0,
			base:      2,
			expDigits: 1,
		},
		{
			ID:        testhelper.MkID("v: 0, base 8"),
			v:         0,
			base:      8,
			expDigits: 1,
		},
		{
			ID:        testhelper.MkID("v: 0, base 10"),
			v:         0,
			base:      10,
			expDigits: 1,
		},
		{
			ID:        testhelper.MkID("v: 0, base 16"),
			v:         0,
			base:      16,
			expDigits: 1,
		},
		{
			ID:        testhelper.MkID("v: 8, base 2"),
			v:         8,
			base:      2,
			expDigits: 4,
		},
		{
			ID:        testhelper.MkID("v: 8, base 8"),
			v:         8,
			base:      8,
			expDigits: 2,
		},
		{
			ID:        testhelper.MkID("v: 8, base 10"),
			v:         8,
			base:      10,
			expDigits: 1,
		},
	}

	for _, tc := range testCases {
		panicked, panicVal := testhelper.PanicSafe(func() {
			digits := mathutil.DigitsInBase(tc.v, tc.base)
			testhelper.DiffInt(t, tc.IDStr(), "digits", digits, tc.expDigits)

			if tc.v > 0 {
				digits = mathutil.DigitsInBaseUnsigned(uint64(tc.v), tc.base)
				testhelper.DiffInt(
					t, tc.IDStr(), "digits", digits, tc.expDigits)

				digits = mathutil.DigitsInBase(tc.v*-1.0, tc.base)
				testhelper.DiffInt(
					t, tc.IDStr(), "digits", digits, tc.expDigits+1)
			}
		})
		testhelper.CheckExpPanic(t, panicked, panicVal, tc)
	}
}
