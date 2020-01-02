package slices

//go:generate genny -in $GOFILE -out ${GOPACKAGE}_genny.go gen "Some=BUILTINS"

import "github.com/cheekybits/genny/generic"

type Some generic.Type

type OfSome []Some
type OfSomeSt struct {
	*OfSome
}

type OfSomeIter []Some

func (__ OfSome) At(i int) Some {
	return __[i]
}

func (__ OfSome) Count() int {
	return len(__)
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
