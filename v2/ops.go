package slices

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Range[N Number](start, end, step N) []N {
	l := end - start
	if l < 0 {
		l = 0
	}
	res := make([]N, 0, int(l/step))
	var i N
	for i = 0; i < l; i += step {
		res = append(res, start+i)
	}
	return res
}
