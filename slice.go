package slices

//go:generate genny -in $GOFILE -out ${GOPACKAGE}_genny.go gen "Some=BUILTINS"

import "github.com/cheekybits/genny/generic"

type Some generic.Type

type OfSomeIf interface {
	Get(int) Some
	Set(int, Some) Some
	Len() int
}

type OfSomeIf32 interface {
	Get(int32) Some
	Set(int32, Some) Some
	Len() int32
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

type OfSomeSt struct {
	somes OfSome
}

func NewOfSomeSt(i int) *OfSomeSt {
	return &OfSomeSt{somes: OfSome(make([]Some, i))}
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

type OfSomeIter []Some

func (__ OfSomeIter) Range(f func(i int, d Some) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfSomeIter) Map(f func(i int, d Some) Some) OfSomeIter {
	rval := make([]Some, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfSomeIter(rval)
}
