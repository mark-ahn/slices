package slices

//go:generate genny -in $GOFILE -out slice_of_template_genny.go gen "Some=interface{}"

// import "github.com/cheekybits/genny/generic"
// type Some generic.Type

type SliceOfSomeIterIf interface {
	Range(f func(i int, d Some) bool)
	Map(f func(i int, d Some) Some) SliceOfSomeMutIf
}
type SliceOfSomeIf interface {
	Get(int) Some
	Len() int
}
type SliceOfSomeMutIf interface {
	SliceOfSomeIf
	Set(int, Some) Some
}

type SliceOfSomeAsIterIf interface {
	AsIter() SliceOfSomeIterIf
}

type SliceOfSomeIf32 interface {
	Get(int32) Some
	Len() int32
}
type SliceOfSomeIfMut32 interface {
	SliceOfSomeIf32
	Set(int32, Some) Some
}

func SliceOfSomeInto(__ SliceOfSomeIf) []Some {
	switch d := __.(type) {
	case SliceOfSome:
		return []Some(d)
	case *SliceOfSomeSt:
		return []Some(d.somes)
	case nil:
		return nil
	default:
		res := make([]Some, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
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

func (__ SliceOfSome) AsIter() SliceOfSomeIterIf {
	return SliceOfSomeIter(__)
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

func (__ SliceOfSomeI32) AsIter() SliceOfSomeIterIf {
	return SliceOfSomeIter(__)
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

func (__ *SliceOfSomeSt) AsIter() SliceOfSomeIterIf {
	return __.somes.AsIter()
}

type SliceOfSomeIter []Some

func (__ SliceOfSomeIter) Range(f func(i int, d Some) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ SliceOfSomeIter) Map(f func(i int, d Some) Some) SliceOfSomeMutIf {
	rval := make([]Some, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return SliceOfSome(rval)
}
