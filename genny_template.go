package slices

import "github.com/cheekybits/genny/generic"

type Some generic.Type

type _Prefix_SomeIf interface {
	Get(int) Some
	Len() int
}
type _Prefix_SomeMutIf interface {
	_Prefix_SomeIf
	Set(int, Some) Some
}
type _Prefix_SomeAsIterIf interface {
	AsIter() _Prefix_SomeIterIf
}

type _Prefix_SomeIterIf interface {
	Range(f func(i int, d Some) bool)
	Map(f func(i int, d Some) Some) _Prefix_SomeMutIf
}

func _Prefix_SomeInto(__ _Prefix_SomeIf) []Some {
	switch d := __.(type) {
	case _Prefix_Some:
		return []Some(d)
	case *_Prefix_SomeSt:
		return []Some(d.__Prefix_Some)
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

type _Prefix_Some []Some

func (__ _Prefix_Some) Get(i int) Some {
	return __[i]
}
func (__ _Prefix_Some) Set(i int, d Some) Some {
	old := __[i]
	__[i] = d
	return old
}

func (__ _Prefix_Some) Len() int {
	return len(__)
}

func (__ _Prefix_Some) AsIter() _Prefix_SomeIterIf {
	return _Prefix_SomeIter(__)
}

type __Prefix_Some = _Prefix_Some
type _Prefix_SomeSt struct {
	// do not want to export but want to use embedding method
	__Prefix_Some
}

func New_Prefix_SomeStFrom(list []Some) *_Prefix_SomeSt {
	return &_Prefix_SomeSt{__Prefix_Some: _Prefix_Some(list)}
}
func New_Prefix_SomeSt(i int) *_Prefix_SomeSt {
	return New_Prefix_SomeStFrom(make([]Some, i))
}

type _Prefix_SomeIter []Some

func (__ _Prefix_SomeIter) Range(f func(i int, d Some) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ _Prefix_SomeIter) Map(f func(i int, d Some) Some) _Prefix_SomeMutIf {
	rval := make([]Some, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return _Prefix_Some(rval)
}
