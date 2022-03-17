package slices_test

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"

// 	"github.com/mark-ahn/slices/v2"
// )

// func TestNil(t *testing.T) {
// 	res := slices.Into[string](nil)
// 	if res != nil {
// 		t.Errorf("expect nil, got %v", res)
// 	}
// }
// func TestT(t *testing.T) {
// 	ss := slices.Of[string]([]string{
// 		"0", "1", "2", "3",
// 	})

// 	if "3" != ss.Get(3) {
// 		t.Errorf("expect %v, got %v", "3", ss.Get(3))
// 	}
// }

// func TestInto(t *testing.T) {
// 	expect := []string{
// 		"0", "1", "2", "3",
// 	}
// 	ss := slices.Of[string](expect)

// 	got := slices.Into[string](ss)
// 	if !reflect.DeepEqual(expect, got) {
// 		t.Errorf("expect %v:%T, got %v:%T", expect, expect, got, got)
// 	}
// }

// func TestTSt(t *testing.T) {
// 	var ss interface {
// 		slices.MutIf[string]
// 		slices.AsIterIf[string]
// 	}
// 	// ss = slices.OfString([]string{
// 	// 	"0", "1", "2", "3",
// 	// })
// 	ss = slices.NewSliceOfSt[string](4)
// 	ss.Set(0, "0")
// 	ss.Set(1, "1")
// 	ss.Set(2, "2")
// 	ss.Set(3, "3")

// 	ss.Set(3, "30")

// 	ss.AsIter().Range(slices.LoopFunc[string](func(_ int, d string) bool {
// 		fmt.Println(d)
// 		return true
// 	}))

// 	if "30" != ss.Get(3) {
// 		t.Errorf("expect %v, got %v", "30", ss.Get(3))
// 	}
// }

// func TestSum(t *testing.T) {
// 	is := slices.Iter[int]([]int{
// 		1, 2, 3, 4, 5,
// 	})

// 	sum := 0
// 	is.Range(slices.LoopFunc[int](func(_ int, d int) bool {
// 		sum += d
// 		return true
// 	}))
// 	if sum != 15 {
// 		t.Errorf("expect %v, got %v", 15, sum)
// 	}

// 	res := is.Map(slices.MapFunc[int](func(_ int, d int) int { return d * 2 }))
// 	expect := slices.Iter[int]([]int{2, 4, 6, 8, 10}).Map(slices.MapFunc[int](func(_ int, d int) int { return d }))
// 	if !reflect.DeepEqual(expect, res) {
// 		t.Errorf("expect %v:%T, got %v:%T", expect, expect, res, res)
// 	}
// }
