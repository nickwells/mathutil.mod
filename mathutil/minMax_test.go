package mathutil_test

import (
	"testing"

	"github.com/nickwells/mathutil.mod/mathutil"
	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestMinMaxOf(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		vals     []float64
		expMin   float64
		expMax   float64
		panicExp bool
	}{
		{
			ID:     testhelper.MkID("all good - one val"),
			vals:   []float64{1.0},
			expMin: 1.0,
			expMax: 1.0,
		},
		{
			ID:     testhelper.MkID("all good - multiple vals"),
			vals:   []float64{-1.0, 2.0, 3.0},
			expMin: -1.0,
			expMax: 3.0,
		},
		{
			ID:     testhelper.MkID("all good - multiple vals, out of order"),
			vals:   []float64{2.0, 3.0, -1.0},
			expMin: -1.0,
			expMax: 3.0,
		},
		{
			ID:       testhelper.MkID("panic expected"),
			vals:     []float64{},
			panicExp: true,
		},
	}

	for _, tc := range testCases {
		v, panicked, panicVal := panicSafeFloat64(mathutil.MinOf, tc.vals)
		if panicOK(t, tc.IDStr()+": MinOf", panicked, tc.panicExp, panicVal) {
			testhelper.DiffFloat64(t, tc.IDStr(), "min", v, tc.expMin, 0.0)
		}

		v, panicked, panicVal = panicSafeFloat64(mathutil.MaxOf, tc.vals)
		if panicOK(t, tc.IDStr()+": MaxOf", panicked, tc.panicExp, panicVal) {
			testhelper.DiffFloat64(t, tc.IDStr(), "max", v, tc.expMax, 0.0)
		}
	}

}

func TestMinMaxOfInt(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		vals     []int
		panicExp bool
		expMin   int
		expMax   int
	}{
		{
			ID:     testhelper.MkID("all good - one val"),
			vals:   []int{1},
			expMin: 1,
			expMax: 1,
		},
		{
			ID:     testhelper.MkID("all good - multiple vals"),
			vals:   []int{-1, 2, 3},
			expMin: -1,
			expMax: 3,
		},
		{
			ID:     testhelper.MkID("all good - multiple vals, out of order"),
			vals:   []int{2, 3, -1},
			expMin: -1,
			expMax: 3,
		},
		{
			ID:       testhelper.MkID("panic expected"),
			vals:     []int{},
			panicExp: true,
		},
	}

	for _, tc := range testCases {
		v, panicked, panicVal := panicSafeInt(mathutil.MinOfInt, tc.vals)
		if panicOK(t, tc.IDStr()+": MinOfInt", panicked, tc.panicExp, panicVal) {
			testhelper.DiffInt(t, tc.IDStr(), "min", v, tc.expMin)
		}

		v, panicked, panicVal = panicSafeInt(mathutil.MaxOfInt, tc.vals)
		if panicOK(t, tc.IDStr()+": MaxOfInt", panicked, tc.panicExp, panicVal) {
			testhelper.DiffInt(t, tc.IDStr(), "max", v, tc.expMax)
		}
	}

}

// panicOK will check that the panic was as expected and return true if all
// was as expected, otherwise it will report the problem and return false
func panicOK(t *testing.T, name string, panicked, expected bool, pVal interface{}) bool {
	t.Helper()

	if panicked && !expected {
		t.Log(name)
		t.Errorf("\t: unexpected panic: %v", pVal)
		return false
	}

	if !panicked && expected {
		t.Log(name)
		t.Errorf("\t: panic expected but not seen")
		return false
	}
	return true
}

// panicSafeFloat64 calls the passed function with the supplied vals, it
// recovers from any panic and returns whether or not a panic was detected
// and the panic value
func panicSafeFloat64(f func(...float64) float64, vals []float64) (result float64, panicked bool, panicVal interface{}) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			panicVal = r
		}
	}()

	result = f(vals...)
	return
}

// panicSafeInt calls the passed function with the supplied vals, it
// recovers from any panic and returns whether or not a panic was detected
// and the panic value
func panicSafeInt(f func(...int) int, vals []int) (result int, panicked bool, panicVal interface{}) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			panicVal = r
		}
	}()

	result = f(vals...)
	return
}
