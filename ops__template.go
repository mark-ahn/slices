package slices

type _Some struct{}

func (_ _Some) Range(start, end int) []Some {
	l := end - start
	if l < 0 {
		l = 0
	}
	res := make([]Some, l)
	for i := 0; i < l; i += 1 {
		res[i] = Some(start + i)
	}
	return res
}

var SomeT _Some
