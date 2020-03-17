package slices

import "github.com/cheekybits/genny/generic"

//go:generate genny -in genny_template.go -out genny_template__gen.go gen "_Prefix_=Of Some=BUILTINS,interface{}"
//go:generate genny -in ops__template.go -out ops__template__gen.go gen "_Prefix_=Of Some=NUMBERS"

type Some generic.Type
