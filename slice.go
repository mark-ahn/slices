package slices

//go:generate genny -in $GOFILE -out ${GOPACKAGE}_genny.go gen "Some=BUILTINS"

import "github.com/cheekybits/genny/generic"

type Some generic.Type

type OfSome []Some

func (__ OfSome) At(i int) Some {
	return __[i]
}

func (__ OfSome) Count() int {
	return len(__)
}
