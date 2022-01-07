package slices

type If[Some any] interface {
	Get(int) Some
	Len() int
}
type MutIf[Some any] interface {
	If[Some]
	Set(int, Some) Some
}
type AsIterIf[Some any] interface {
	AsIter() IterIf[Some]
}

type Looper[Some any] interface {
	LoopItem(i int, d Some) bool
}
type Mapper[Some any] interface {
	MapItem(i int, d Some) Some
}
type LoopFunc[Some any] func(i int, d Some) bool

func (__ LoopFunc[Some]) LoopItem(i int, d Some) bool {
	return __(i, d)
}

type MapFunc[Some any] func(i int, d Some) Some

func (__ MapFunc[Some]) MapItem(i int, d Some) Some {
	return __(i, d)
}

type IterIf[Some any] interface {
	Range(fntr Looper[Some])
	Map(fntr Mapper[Some]) MutIf[Some]
}

func Into[Some any](__ If[Some]) []Some {
	switch d := __.(type) {
	case SliceOf[Some]:
		return []Some(d)
	case *SliceOfSt[Some]:
		return []Some(d.SliceOf)
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

type SliceOf[Some any] []Some

func (__ SliceOf[Some]) Get(i int) Some {
	return __[i]
}
func (__ SliceOf[Some]) Set(i int, d Some) Some {
	old := __[i]
	__[i] = d
	return old
}

func (__ SliceOf[Some]) Len() int {
	return len(__)
}

func (__ SliceOf[Some]) AsIter() IterIf[Some] {
	return Iter[Some](__)
}

// type sliceOf[Some any] = SliceOf[Some]
type SliceOfSt[Some any] struct {
	// do not want to export but want to use embedding method
	SliceOf[Some]
}

func NewSliceOfStFrom[Some any](list []Some) *SliceOfSt[Some] {
	return &SliceOfSt[Some]{SliceOf: (SliceOf[Some](list))}
}
func NewSliceOfSt[Some any](i int) *SliceOfSt[Some] {
	return NewSliceOfStFrom[Some](make([]Some, i))
}

type Iter[Some any] []Some

func (__ Iter[Some]) Range(fntr Looper[Some]) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}

func (__ Iter[Some]) Map(fntr Mapper[Some]) MutIf[Some] {
	rval := make([]Some, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return SliceOf[Some](rval)
}
