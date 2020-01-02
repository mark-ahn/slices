// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slices

type OfBool []bool
type OfBoolIter []bool

func (__ OfBool) At(i int) bool {
	return __[i]
}

func (__ OfBool) Count() int {
	return len(__)
}

func (__ OfBoolIter) Range(f func(i int, d bool) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfBoolIter) Map(f func(i int, d bool) bool) OfBoolIter {
	rval := make([]bool, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfBoolIter(rval)
}

type OfByte []byte
type OfByteIter []byte

func (__ OfByte) At(i int) byte {
	return __[i]
}

func (__ OfByte) Count() int {
	return len(__)
}

func (__ OfByteIter) Range(f func(i int, d byte) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfByteIter) Map(f func(i int, d byte) byte) OfByteIter {
	rval := make([]byte, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfByteIter(rval)
}

type OfComplex128 []complex128
type OfComplex128Iter []complex128

func (__ OfComplex128) At(i int) complex128 {
	return __[i]
}

func (__ OfComplex128) Count() int {
	return len(__)
}

func (__ OfComplex128Iter) Range(f func(i int, d complex128) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfComplex128Iter) Map(f func(i int, d complex128) complex128) OfComplex128Iter {
	rval := make([]complex128, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfComplex128Iter(rval)
}

type OfComplex64 []complex64
type OfComplex64Iter []complex64

func (__ OfComplex64) At(i int) complex64 {
	return __[i]
}

func (__ OfComplex64) Count() int {
	return len(__)
}

func (__ OfComplex64Iter) Range(f func(i int, d complex64) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfComplex64Iter) Map(f func(i int, d complex64) complex64) OfComplex64Iter {
	rval := make([]complex64, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfComplex64Iter(rval)
}

type OfError []error
type OfErrorIter []error

func (__ OfError) At(i int) error {
	return __[i]
}

func (__ OfError) Count() int {
	return len(__)
}

func (__ OfErrorIter) Range(f func(i int, d error) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfErrorIter) Map(f func(i int, d error) error) OfErrorIter {
	rval := make([]error, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfErrorIter(rval)
}

type OfFloat32 []float32
type OfFloat32Iter []float32

func (__ OfFloat32) At(i int) float32 {
	return __[i]
}

func (__ OfFloat32) Count() int {
	return len(__)
}

func (__ OfFloat32Iter) Range(f func(i int, d float32) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfFloat32Iter) Map(f func(i int, d float32) float32) OfFloat32Iter {
	rval := make([]float32, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfFloat32Iter(rval)
}

type OfFloat64 []float64
type OfFloat64Iter []float64

func (__ OfFloat64) At(i int) float64 {
	return __[i]
}

func (__ OfFloat64) Count() int {
	return len(__)
}

func (__ OfFloat64Iter) Range(f func(i int, d float64) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfFloat64Iter) Map(f func(i int, d float64) float64) OfFloat64Iter {
	rval := make([]float64, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfFloat64Iter(rval)
}

type OfInt []int
type OfIntIter []int

func (__ OfInt) At(i int) int {
	return __[i]
}

func (__ OfInt) Count() int {
	return len(__)
}

func (__ OfIntIter) Range(f func(i int, d int) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfIntIter) Map(f func(i int, d int) int) OfIntIter {
	rval := make([]int, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfIntIter(rval)
}

type OfInt16 []int16
type OfInt16Iter []int16

func (__ OfInt16) At(i int) int16 {
	return __[i]
}

func (__ OfInt16) Count() int {
	return len(__)
}

func (__ OfInt16Iter) Range(f func(i int, d int16) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfInt16Iter) Map(f func(i int, d int16) int16) OfInt16Iter {
	rval := make([]int16, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfInt16Iter(rval)
}

type OfInt32 []int32
type OfInt32Iter []int32

func (__ OfInt32) At(i int) int32 {
	return __[i]
}

func (__ OfInt32) Count() int {
	return len(__)
}

func (__ OfInt32Iter) Range(f func(i int, d int32) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfInt32Iter) Map(f func(i int, d int32) int32) OfInt32Iter {
	rval := make([]int32, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfInt32Iter(rval)
}

type OfInt64 []int64
type OfInt64Iter []int64

func (__ OfInt64) At(i int) int64 {
	return __[i]
}

func (__ OfInt64) Count() int {
	return len(__)
}

func (__ OfInt64Iter) Range(f func(i int, d int64) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfInt64Iter) Map(f func(i int, d int64) int64) OfInt64Iter {
	rval := make([]int64, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfInt64Iter(rval)
}

type OfInt8 []int8
type OfInt8Iter []int8

func (__ OfInt8) At(i int) int8 {
	return __[i]
}

func (__ OfInt8) Count() int {
	return len(__)
}

func (__ OfInt8Iter) Range(f func(i int, d int8) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfInt8Iter) Map(f func(i int, d int8) int8) OfInt8Iter {
	rval := make([]int8, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfInt8Iter(rval)
}

type OfRune []rune
type OfRuneIter []rune

func (__ OfRune) At(i int) rune {
	return __[i]
}

func (__ OfRune) Count() int {
	return len(__)
}

func (__ OfRuneIter) Range(f func(i int, d rune) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfRuneIter) Map(f func(i int, d rune) rune) OfRuneIter {
	rval := make([]rune, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfRuneIter(rval)
}

type OfString []string
type OfStringIter []string

func (__ OfString) At(i int) string {
	return __[i]
}

func (__ OfString) Count() int {
	return len(__)
}

func (__ OfStringIter) Range(f func(i int, d string) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfStringIter) Map(f func(i int, d string) string) OfStringIter {
	rval := make([]string, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfStringIter(rval)
}

type OfUint []uint
type OfUintIter []uint

func (__ OfUint) At(i int) uint {
	return __[i]
}

func (__ OfUint) Count() int {
	return len(__)
}

func (__ OfUintIter) Range(f func(i int, d uint) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfUintIter) Map(f func(i int, d uint) uint) OfUintIter {
	rval := make([]uint, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfUintIter(rval)
}

type OfUint16 []uint16
type OfUint16Iter []uint16

func (__ OfUint16) At(i int) uint16 {
	return __[i]
}

func (__ OfUint16) Count() int {
	return len(__)
}

func (__ OfUint16Iter) Range(f func(i int, d uint16) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfUint16Iter) Map(f func(i int, d uint16) uint16) OfUint16Iter {
	rval := make([]uint16, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfUint16Iter(rval)
}

type OfUint32 []uint32
type OfUint32Iter []uint32

func (__ OfUint32) At(i int) uint32 {
	return __[i]
}

func (__ OfUint32) Count() int {
	return len(__)
}

func (__ OfUint32Iter) Range(f func(i int, d uint32) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfUint32Iter) Map(f func(i int, d uint32) uint32) OfUint32Iter {
	rval := make([]uint32, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfUint32Iter(rval)
}

type OfUint64 []uint64
type OfUint64Iter []uint64

func (__ OfUint64) At(i int) uint64 {
	return __[i]
}

func (__ OfUint64) Count() int {
	return len(__)
}

func (__ OfUint64Iter) Range(f func(i int, d uint64) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfUint64Iter) Map(f func(i int, d uint64) uint64) OfUint64Iter {
	rval := make([]uint64, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfUint64Iter(rval)
}

type OfUint8 []uint8
type OfUint8Iter []uint8

func (__ OfUint8) At(i int) uint8 {
	return __[i]
}

func (__ OfUint8) Count() int {
	return len(__)
}

func (__ OfUint8Iter) Range(f func(i int, d uint8) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfUint8Iter) Map(f func(i int, d uint8) uint8) OfUint8Iter {
	rval := make([]uint8, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfUint8Iter(rval)
}

type OfUintptr []uintptr
type OfUintptrIter []uintptr

func (__ OfUintptr) At(i int) uintptr {
	return __[i]
}

func (__ OfUintptr) Count() int {
	return len(__)
}

func (__ OfUintptrIter) Range(f func(i int, d uintptr) bool) {
	for i := range __ {
		if !f(i, __[i]) {
			break
		}
	}
}
func (__ OfUintptrIter) Map(f func(i int, d uintptr) uintptr) OfUintptrIter {
	rval := make([]uintptr, len(__))
	for i := range __ {
		rval[i] = f(i, __[i])
	}
	return OfUintptrIter(rval)
}
