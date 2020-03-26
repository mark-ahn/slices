package slices

type _Num struct{}

func (_ _Num) Range(start, end, step Num) []Num {
	l := end - start
	if l < 0 {
		l = 0
	}
	res := make([]Num, 0, int(l/step))
	var i Num
	for i = 0; i < l; i += step {
		res = append(res, start+i)
	}
	return res
}

var NumT _Num
