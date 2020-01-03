package slices

//go:generate genny -in $GOFILE -out slice_of_template_genny.go gen "Some=interface{}"

// type Some generic.Type

type SliceOfSomeIf interface {
	Get(int) Some
	Set(int, Some) Some
	Len() int
}

type SliceOfSomeIf32 interface {
	Get(int32) Some
	Set(int32, Some) Some
	Len() int32
}

type SliceOfSome []Some

func (__ SliceOfSome) Get(i int) Some {
	return __[i]
}
func (__ SliceOfSome) Set(i int, d Some) Some {
	old := __[i]
	__[i] = d
	return old
}

func (__ SliceOfSome) Len() int {
	return len(__)
}

type SliceOfSomeI32 []Some

func (__ SliceOfSomeI32) Get(i int32) Some {
	return __[int(i)]
}

func (__ SliceOfSomeI32) Set(i int32, d Some) Some {
	old := __[i]
	__[i] = d
	return old
}

func (__ SliceOfSomeI32) Len() int32 {
	return int32(len(__))
}

type SliceOfSomeSt struct {
	somes SliceOfSome
}

func NewSliceOfSomeSt(i int) *SliceOfSomeSt {
	return &SliceOfSomeSt{somes: SliceOfSome(make([]Some, i))}
}

func (__ *SliceOfSomeSt) Get(i int) Some {
	return __.somes.Get(i)
}
func (__ *SliceOfSomeSt) Set(i int, d Some) Some {
	return __.somes.Set(i, d)
}

func (__ *SliceOfSomeSt) Len() int {
	return __.somes.Len()
}

type SliceOfSomeIter []Some

func (__ SliceOfSomeIter) Range(f func(i int, d Some) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ SliceOfSomeIter) Map(f func(i int, d Some) Some) SliceOfSomeIter {
	rval := make([]Some, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return SliceOfSomeIter(rval)
}
