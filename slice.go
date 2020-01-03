package slices

//go:generate genny -in $GOFILE -out ${GOPACKAGE}_genny.go gen "Some=BUILTINS"

import "github.com/cheekybits/genny/generic"

type Some generic.Type

type OfSomeInf interface {
	At(int) Some
	Count() int
}

type OfSomeInf32 interface {
	At(int32) Some
	Count() int32
}

type OfSome []Some
type OfSomeI32 []Some

type OfSomeIter []Some

func (__ OfSome) At(i int) Some {
	return __[i]
}

func (__ OfSome) Count() int {
	return len(__)
}

func (__ OfSomeI32) At(i int32) Some {
	return __[int(i)]
}

func (__ OfSomeI32) Count() int32 {
	return int32(len(__))
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
