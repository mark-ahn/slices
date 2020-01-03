package slices

//go:generate genny -in $GOFILE -out slice_of_template_genny.go gen "SomeT=interface{}"

import "github.com/cheekybits/genny/generic"

type SomeT generic.Type

type SomeTSliceInf interface {
	At(int) SomeT
	Count() int
}

type SomeTSliceInf32 interface {
	At(int32) SomeT
	Count() int32
}

type SomeTSlice []SomeT
type SomeTSliceI32 []SomeT

type SomeTSliceIter []SomeT

func (__ SomeTSlice) At(i int) SomeT {
	return __[i]
}

func (__ SomeTSlice) Count() int {
	return len(__)
}

func (__ SomeTSliceI32) At(i int32) SomeT {
	return __[int(i)]
}

func (__ SomeTSliceI32) Count() int32 {
	return int32(len(__))
}

func (__ SomeTSliceIter) Range(f func(i int, d SomeT) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ SomeTSliceIter) Map(f func(i int, d SomeT) SomeT) SomeTSliceIter {
	rval := make([]SomeT, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return SomeTSliceIter(rval)
}
