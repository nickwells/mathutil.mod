package mathutil_test

import (
	"testing"

	"github.com/nickwells/mathutil.mod/v2/mathutil"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestBitsInType(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		v       any
		expBits int
	}{
		{
			ID:      testhelper.MkID("int64"),
			v:       int64(42),
			expBits: 64,
		},
		{
			ID:      testhelper.MkID("uint64"),
			v:       uint64(42),
			expBits: 64,
		},
		{
			ID:      testhelper.MkID("int8"),
			v:       int8(42),
			expBits: 8,
		},
		{
			ID:      testhelper.MkID("uint8"),
			v:       uint8(42),
			expBits: 8,
		},
		{
			ID:      testhelper.MkID("float64"),
			v:       float64(42),
			expBits: 64,
		},
		{
			ID:      testhelper.MkID("float32"),
			v:       float32(42),
			expBits: 32,
		},
		{
			ID:      testhelper.MkID("string"),
			v:       "hello",
			expBits: 128,
		},
		{
			ID:      testhelper.MkID("struct"),
			v:       struct{ a, b, c, d int8 }{4, 2, 1, 0},
			expBits: 8 * 4,
		},
	}

	for _, tc := range testCases {
		b := mathutil.BitsInType(tc.v)
		testhelper.DiffInt(t, tc.IDStr(), "bits", b, tc.expBits)
	}
}
