package slices

type numbers interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Range[Num numbers](start, end, step Num) []Num {
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
