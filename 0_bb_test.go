package slices_test

import (
	"reflect"
	"testing"

	"github.com/mark-ahn/slices"
)

func TestT(t *testing.T) {
	ss := slices.OfString([]string{
		"0", "1", "2", "3",
	})

	if "3" != ss.At(3) {
		t.Errorf("expect %v, got %v", "3", ss.At(3))
	}
}

func TestTSt(t *testing.T) {
	var ss slices.OfStringInf
	ss = slices.OfString([]string{
		"0", "1", "2", "3",
	})

	if "3" != ss.At(3) {
		t.Errorf("expect %v, got %v", "3", ss.At(3))
	}
}

func TestSum(t *testing.T) {
	is := slices.OfIntIter([]int{
		1, 2, 3, 4, 5,
	})

	sum := 0
	is.Range(func(_ int, d int) bool {
		sum += d
		return true
	})
	if sum != 15 {
		t.Errorf("expect %v, got %v", 15, sum)
	}

	res := is.Map(func(_ int, d int) int { return d * 2 })
	expect := slices.OfIntIter([]int{2, 4, 6, 8, 10})
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("expect %v:%T, got %v:%T", expect, expect, res, res)
	}
}
