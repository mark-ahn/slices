package slices

//go:generate genny -in $GOFILE -out ${GOPACKAGE}_genny.go gen "Some=BUILTINS"

import "github.com/cheekybits/genny/generic"

type Some generic.Type

type OfSomeIterIf interface {
	Range(f func(i int, d Some) bool)
	Map(f func(i int, d Some) Some) OfSomeMutIf
}
type OfSomeIf interface {
	Get(int) Some
	Len() int
}
type OfSomeMutIf interface {
	OfSomeIf
	Set(int, Some) Some
}

type OfSomeAsIterIf interface {
	AsIter() OfSomeIterIf
}

type OfSomeI32If interface {
	Get(int32) Some
	Len() int32
}
type OfSomeI32MutIf interface {
	OfSomeI32If
	Set(int32, Some) Some
}

func OfSomeInto(__ OfSomeIf) []Some {
	switch d := __.(type) {
	case OfSome:
		return []Some(d)
	case *OfSomeSt:
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

type OfSome []Some

func (__ OfSome) Get(i int) Some {
	return __[i]
}
func (__ OfSome) Set(i int, d Some) Some {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfSome) Len() int {
	return len(__)
}

func (__ OfSome) AsIter() OfSomeIterIf {
	return OfSomeIter(__)
}

func OfSomeI32Into(__ OfSomeI32If) []Some {
	switch d := __.(type) {
	case OfSomeI32:
		return []Some(d)
	case *OfSomeStI32:
		return []Some(d.somes)
	case nil:
		return nil
	default:
		res := make([]Some, __.Len())
		for i := int32(0); i < int32(len(res)); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfSomeI32 []Some

func (__ OfSomeI32) Get(i int32) Some {
	return __[int(i)]
}

func (__ OfSomeI32) Set(i int32, d Some) Some {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfSomeI32) Len() int32 {
	return int32(len(__))
}

func (__ OfSomeI32) AsIter() OfSomeIterIf {
	return OfSomeIter(__)
}

type OfSomeSt struct {
	somes OfSome
}

func NewOfSomeSt(i int) *OfSomeSt {
	return &OfSomeSt{somes: OfSome(make([]Some, i))}
}
func NewOfSomeStFrom(list []Some) *OfSomeSt {
	return &OfSomeSt{somes: OfSome(list)}
}

func (__ *OfSomeSt) Get(i int) Some {
	return __.somes.Get(i)
}
func (__ *OfSomeSt) Set(i int, d Some) Some {
	return __.somes.Set(i, d)
}

func (__ *OfSomeSt) Len() int {
	return __.somes.Len()
}

func (__ *OfSomeSt) AsIter() OfSomeIterIf {
	return __.somes.AsIter()
}

type OfSomeStI32 struct {
	somes OfSomeI32
}

func NewOfSomeStI32(i int32) *OfSomeStI32 {
	return &OfSomeStI32{somes: OfSomeI32(make([]Some, i))}
}

func NewOfSomeStI32From(list []Some) *OfSomeStI32 {
	return &OfSomeStI32{somes: OfSomeI32(list)}
}

func (__ *OfSomeStI32) Get(i int32) Some {
	return __.somes.Get(i)
}
func (__ *OfSomeStI32) Set(i int32, d Some) Some {
	return __.somes.Set(i, d)
}

func (__ *OfSomeStI32) Len() int32 {
	return __.somes.Len()
}

func (__ *OfSomeStI32) AsIter() OfSomeIterIf {
	return __.somes.AsIter()
}

type OfSomeIter []Some

func (__ OfSomeIter) Range(f func(i int, d Some) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfSomeIter) Map(f func(i int, d Some) Some) OfSomeMutIf {
	rval := make([]Some, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfSome(rval)
}
