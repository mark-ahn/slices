package slices

//go:generate genny -in $GOFILE -out ${GOPACKAGE}_genny.go gen "Some=BUILTINS"

import "github.com/cheekybits/genny/generic"

type Some generic.Type

type OfSomeInf interface {
	At(int) Some
	Count() int
}

type OfSome []Some
type OfSomeSt struct {
	in []Some
}

func NewSomeSt(d []Some) *OfSomeSt {
	return &OfSomeSt{
		in: d,
	}
}

type OfSomeIter []Some

func (__ OfSome) At(i int) Some {
	return __[i]
}

func (__ OfSome) Count() int {
	return len(__)
}

func (__ *OfSomeSt) At(i int) Some {
	return __.in[i]
}

func (__ *OfSomeSt) Count() int {
	return len(__.in)
}

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
