package slices

//go:generate genny -in $GOFILE -out slice_of_template_genny.go gen "SomeT=interface{}"

import "github.com/cheekybits/genny/generic"

type SomeT generic.Type

type SomeTSliceInf interface {
	At(int) Some
	Set(int, Some)
	Count() int
}

type SomeTSliceInf32 interface {
	At(int32) Some
	Set(int32, Some)
	Count() int32
}

type SomeTSlice []Some
type SomeTSliceI32 []Some

type SomeTSliceIter []Some

func NewSomeTSlice(i int) SomeTSlice {
	return SomeTSlice(make([]Some, i))
}

func (__ SomeTSlice) At(i int) Some {
	return __[i]
}
func (__ SomeTSlice) Set(i int, d Some) {
	__[i] = d
}

func (__ SomeTSlice) Count() int {
	return len(__)
}

func NewSomeTSliceI32(i int) SomeTSliceI32 {
	return SomeTSliceI32(make([]Some, i))
}

func (__ SomeTSliceI32) At(i int32) Some {
	return __[int(i)]
}

func (__ SomeTSliceI32) Set(i int32, d Some) {
	__[i] = d
}

func (__ SomeTSliceI32) Count() int32 {
	return int32(len(__))
}

func (__ SomeTSliceIter) Range(f func(i int, d Some) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ SomeTSliceIter) Map(f func(i int, d Some) Some) SomeTSliceIter {
	rval := make([]Some, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return SomeTSliceIter(rval)
}