// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slices

type OfBoolIf interface {
	Get(int) bool
	Len() int
}
type OfBoolMutIf interface {
	OfBoolIf
	Set(int, bool) bool
}
type OfBoolAsIterIf interface {
	AsIter() OfBoolIterIf
}

type OfBoolLooper interface {
	LoopItem(i int, d bool) bool
}
type OfBoolMapper interface {
	MapItem(i int, d bool) bool
}
type OfBoolLoopFunc func(i int, d bool) bool

func (__ OfBoolLoopFunc) LoopItem(i int, d bool) bool {
	return __(i, d)
}

type OfBoolMapFunc func(i int, d bool) bool

func (__ OfBoolMapFunc) MapItem(i int, d bool) bool {
	return __(i, d)
}

type OfBoolIterIf interface {
	Range(fntr OfBoolLooper)
	Map(fntr OfBoolMapper) OfBoolMutIf
}

func OfBoolInto(__ OfBoolIf) []bool {
	switch d := __.(type) {
	case OfBool:
		return []bool(d)
	case *OfBoolSt:
		return []bool(d._OfBool)
	case nil:
		return nil
	default:
		res := make([]bool, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfBool []bool

func (__ OfBool) Get(i int) bool {
	return __[i]
}
func (__ OfBool) Set(i int, d bool) bool {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfBool) Len() int {
	return len(__)
}

func (__ OfBool) AsIter() OfBoolIterIf {
	return OfBoolIter(__)
}

type _OfBool = OfBool
type OfBoolSt struct {
	// do not want to export but want to use embedding method
	_OfBool
}

func NewOfBoolStFrom(list []bool) *OfBoolSt {
	return &OfBoolSt{_OfBool: OfBool(list)}
}
func NewOfBoolSt(i int) *OfBoolSt {
	return NewOfBoolStFrom(make([]bool, i))
}

type OfBoolIter []bool

func (__ OfBoolIter) Range(fntr OfBoolLooper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfBoolIter) Map(fntr OfBoolMapper) OfBoolMutIf {
	rval := make([]bool, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfBool(rval)
}

type OfByteIf interface {
	Get(int) byte
	Len() int
}
type OfByteMutIf interface {
	OfByteIf
	Set(int, byte) byte
}
type OfByteAsIterIf interface {
	AsIter() OfByteIterIf
}

type OfByteLooper interface {
	LoopItem(i int, d byte) bool
}
type OfByteMapper interface {
	MapItem(i int, d byte) byte
}
type OfByteLoopFunc func(i int, d byte) bool

func (__ OfByteLoopFunc) LoopItem(i int, d byte) bool {
	return __(i, d)
}

type OfByteMapFunc func(i int, d byte) byte

func (__ OfByteMapFunc) MapItem(i int, d byte) byte {
	return __(i, d)
}

type OfByteIterIf interface {
	Range(fntr OfByteLooper)
	Map(fntr OfByteMapper) OfByteMutIf
}

func OfByteInto(__ OfByteIf) []byte {
	switch d := __.(type) {
	case OfByte:
		return []byte(d)
	case *OfByteSt:
		return []byte(d._OfByte)
	case nil:
		return nil
	default:
		res := make([]byte, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfByte []byte

func (__ OfByte) Get(i int) byte {
	return __[i]
}
func (__ OfByte) Set(i int, d byte) byte {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfByte) Len() int {
	return len(__)
}

func (__ OfByte) AsIter() OfByteIterIf {
	return OfByteIter(__)
}

type _OfByte = OfByte
type OfByteSt struct {
	// do not want to export but want to use embedding method
	_OfByte
}

func NewOfByteStFrom(list []byte) *OfByteSt {
	return &OfByteSt{_OfByte: OfByte(list)}
}
func NewOfByteSt(i int) *OfByteSt {
	return NewOfByteStFrom(make([]byte, i))
}

type OfByteIter []byte

func (__ OfByteIter) Range(fntr OfByteLooper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfByteIter) Map(fntr OfByteMapper) OfByteMutIf {
	rval := make([]byte, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfByte(rval)
}

type OfComplex128If interface {
	Get(int) complex128
	Len() int
}
type OfComplex128MutIf interface {
	OfComplex128If
	Set(int, complex128) complex128
}
type OfComplex128AsIterIf interface {
	AsIter() OfComplex128IterIf
}

type OfComplex128Looper interface {
	LoopItem(i int, d complex128) bool
}
type OfComplex128Mapper interface {
	MapItem(i int, d complex128) complex128
}
type OfComplex128LoopFunc func(i int, d complex128) bool

func (__ OfComplex128LoopFunc) LoopItem(i int, d complex128) bool {
	return __(i, d)
}

type OfComplex128MapFunc func(i int, d complex128) complex128

func (__ OfComplex128MapFunc) MapItem(i int, d complex128) complex128 {
	return __(i, d)
}

type OfComplex128IterIf interface {
	Range(fntr OfComplex128Looper)
	Map(fntr OfComplex128Mapper) OfComplex128MutIf
}

func OfComplex128Into(__ OfComplex128If) []complex128 {
	switch d := __.(type) {
	case OfComplex128:
		return []complex128(d)
	case *OfComplex128St:
		return []complex128(d._OfComplex128)
	case nil:
		return nil
	default:
		res := make([]complex128, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfComplex128 []complex128

func (__ OfComplex128) Get(i int) complex128 {
	return __[i]
}
func (__ OfComplex128) Set(i int, d complex128) complex128 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfComplex128) Len() int {
	return len(__)
}

func (__ OfComplex128) AsIter() OfComplex128IterIf {
	return OfComplex128Iter(__)
}

type _OfComplex128 = OfComplex128
type OfComplex128St struct {
	// do not want to export but want to use embedding method
	_OfComplex128
}

func NewOfComplex128StFrom(list []complex128) *OfComplex128St {
	return &OfComplex128St{_OfComplex128: OfComplex128(list)}
}
func NewOfComplex128St(i int) *OfComplex128St {
	return NewOfComplex128StFrom(make([]complex128, i))
}

type OfComplex128Iter []complex128

func (__ OfComplex128Iter) Range(fntr OfComplex128Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfComplex128Iter) Map(fntr OfComplex128Mapper) OfComplex128MutIf {
	rval := make([]complex128, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfComplex128(rval)
}

type OfComplex64If interface {
	Get(int) complex64
	Len() int
}
type OfComplex64MutIf interface {
	OfComplex64If
	Set(int, complex64) complex64
}
type OfComplex64AsIterIf interface {
	AsIter() OfComplex64IterIf
}

type OfComplex64Looper interface {
	LoopItem(i int, d complex64) bool
}
type OfComplex64Mapper interface {
	MapItem(i int, d complex64) complex64
}
type OfComplex64LoopFunc func(i int, d complex64) bool

func (__ OfComplex64LoopFunc) LoopItem(i int, d complex64) bool {
	return __(i, d)
}

type OfComplex64MapFunc func(i int, d complex64) complex64

func (__ OfComplex64MapFunc) MapItem(i int, d complex64) complex64 {
	return __(i, d)
}

type OfComplex64IterIf interface {
	Range(fntr OfComplex64Looper)
	Map(fntr OfComplex64Mapper) OfComplex64MutIf
}

func OfComplex64Into(__ OfComplex64If) []complex64 {
	switch d := __.(type) {
	case OfComplex64:
		return []complex64(d)
	case *OfComplex64St:
		return []complex64(d._OfComplex64)
	case nil:
		return nil
	default:
		res := make([]complex64, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfComplex64 []complex64

func (__ OfComplex64) Get(i int) complex64 {
	return __[i]
}
func (__ OfComplex64) Set(i int, d complex64) complex64 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfComplex64) Len() int {
	return len(__)
}

func (__ OfComplex64) AsIter() OfComplex64IterIf {
	return OfComplex64Iter(__)
}

type _OfComplex64 = OfComplex64
type OfComplex64St struct {
	// do not want to export but want to use embedding method
	_OfComplex64
}

func NewOfComplex64StFrom(list []complex64) *OfComplex64St {
	return &OfComplex64St{_OfComplex64: OfComplex64(list)}
}
func NewOfComplex64St(i int) *OfComplex64St {
	return NewOfComplex64StFrom(make([]complex64, i))
}

type OfComplex64Iter []complex64

func (__ OfComplex64Iter) Range(fntr OfComplex64Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfComplex64Iter) Map(fntr OfComplex64Mapper) OfComplex64MutIf {
	rval := make([]complex64, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfComplex64(rval)
}

type OfErrorIf interface {
	Get(int) error
	Len() int
}
type OfErrorMutIf interface {
	OfErrorIf
	Set(int, error) error
}
type OfErrorAsIterIf interface {
	AsIter() OfErrorIterIf
}

type OfErrorLooper interface {
	LoopItem(i int, d error) bool
}
type OfErrorMapper interface {
	MapItem(i int, d error) error
}
type OfErrorLoopFunc func(i int, d error) bool

func (__ OfErrorLoopFunc) LoopItem(i int, d error) bool {
	return __(i, d)
}

type OfErrorMapFunc func(i int, d error) error

func (__ OfErrorMapFunc) MapItem(i int, d error) error {
	return __(i, d)
}

type OfErrorIterIf interface {
	Range(fntr OfErrorLooper)
	Map(fntr OfErrorMapper) OfErrorMutIf
}

func OfErrorInto(__ OfErrorIf) []error {
	switch d := __.(type) {
	case OfError:
		return []error(d)
	case *OfErrorSt:
		return []error(d._OfError)
	case nil:
		return nil
	default:
		res := make([]error, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfError []error

func (__ OfError) Get(i int) error {
	return __[i]
}
func (__ OfError) Set(i int, d error) error {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfError) Len() int {
	return len(__)
}

func (__ OfError) AsIter() OfErrorIterIf {
	return OfErrorIter(__)
}

type _OfError = OfError
type OfErrorSt struct {
	// do not want to export but want to use embedding method
	_OfError
}

func NewOfErrorStFrom(list []error) *OfErrorSt {
	return &OfErrorSt{_OfError: OfError(list)}
}
func NewOfErrorSt(i int) *OfErrorSt {
	return NewOfErrorStFrom(make([]error, i))
}

type OfErrorIter []error

func (__ OfErrorIter) Range(fntr OfErrorLooper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfErrorIter) Map(fntr OfErrorMapper) OfErrorMutIf {
	rval := make([]error, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfError(rval)
}

type OfFloat32If interface {
	Get(int) float32
	Len() int
}
type OfFloat32MutIf interface {
	OfFloat32If
	Set(int, float32) float32
}
type OfFloat32AsIterIf interface {
	AsIter() OfFloat32IterIf
}

type OfFloat32Looper interface {
	LoopItem(i int, d float32) bool
}
type OfFloat32Mapper interface {
	MapItem(i int, d float32) float32
}
type OfFloat32LoopFunc func(i int, d float32) bool

func (__ OfFloat32LoopFunc) LoopItem(i int, d float32) bool {
	return __(i, d)
}

type OfFloat32MapFunc func(i int, d float32) float32

func (__ OfFloat32MapFunc) MapItem(i int, d float32) float32 {
	return __(i, d)
}

type OfFloat32IterIf interface {
	Range(fntr OfFloat32Looper)
	Map(fntr OfFloat32Mapper) OfFloat32MutIf
}

func OfFloat32Into(__ OfFloat32If) []float32 {
	switch d := __.(type) {
	case OfFloat32:
		return []float32(d)
	case *OfFloat32St:
		return []float32(d._OfFloat32)
	case nil:
		return nil
	default:
		res := make([]float32, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfFloat32 []float32

func (__ OfFloat32) Get(i int) float32 {
	return __[i]
}
func (__ OfFloat32) Set(i int, d float32) float32 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfFloat32) Len() int {
	return len(__)
}

func (__ OfFloat32) AsIter() OfFloat32IterIf {
	return OfFloat32Iter(__)
}

type _OfFloat32 = OfFloat32
type OfFloat32St struct {
	// do not want to export but want to use embedding method
	_OfFloat32
}

func NewOfFloat32StFrom(list []float32) *OfFloat32St {
	return &OfFloat32St{_OfFloat32: OfFloat32(list)}
}
func NewOfFloat32St(i int) *OfFloat32St {
	return NewOfFloat32StFrom(make([]float32, i))
}

type OfFloat32Iter []float32

func (__ OfFloat32Iter) Range(fntr OfFloat32Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfFloat32Iter) Map(fntr OfFloat32Mapper) OfFloat32MutIf {
	rval := make([]float32, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfFloat32(rval)
}

type OfFloat64If interface {
	Get(int) float64
	Len() int
}
type OfFloat64MutIf interface {
	OfFloat64If
	Set(int, float64) float64
}
type OfFloat64AsIterIf interface {
	AsIter() OfFloat64IterIf
}

type OfFloat64Looper interface {
	LoopItem(i int, d float64) bool
}
type OfFloat64Mapper interface {
	MapItem(i int, d float64) float64
}
type OfFloat64LoopFunc func(i int, d float64) bool

func (__ OfFloat64LoopFunc) LoopItem(i int, d float64) bool {
	return __(i, d)
}

type OfFloat64MapFunc func(i int, d float64) float64

func (__ OfFloat64MapFunc) MapItem(i int, d float64) float64 {
	return __(i, d)
}

type OfFloat64IterIf interface {
	Range(fntr OfFloat64Looper)
	Map(fntr OfFloat64Mapper) OfFloat64MutIf
}

func OfFloat64Into(__ OfFloat64If) []float64 {
	switch d := __.(type) {
	case OfFloat64:
		return []float64(d)
	case *OfFloat64St:
		return []float64(d._OfFloat64)
	case nil:
		return nil
	default:
		res := make([]float64, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfFloat64 []float64

func (__ OfFloat64) Get(i int) float64 {
	return __[i]
}
func (__ OfFloat64) Set(i int, d float64) float64 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfFloat64) Len() int {
	return len(__)
}

func (__ OfFloat64) AsIter() OfFloat64IterIf {
	return OfFloat64Iter(__)
}

type _OfFloat64 = OfFloat64
type OfFloat64St struct {
	// do not want to export but want to use embedding method
	_OfFloat64
}

func NewOfFloat64StFrom(list []float64) *OfFloat64St {
	return &OfFloat64St{_OfFloat64: OfFloat64(list)}
}
func NewOfFloat64St(i int) *OfFloat64St {
	return NewOfFloat64StFrom(make([]float64, i))
}

type OfFloat64Iter []float64

func (__ OfFloat64Iter) Range(fntr OfFloat64Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfFloat64Iter) Map(fntr OfFloat64Mapper) OfFloat64MutIf {
	rval := make([]float64, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfFloat64(rval)
}

type OfIntIf interface {
	Get(int) int
	Len() int
}
type OfIntMutIf interface {
	OfIntIf
	Set(int, int) int
}
type OfIntAsIterIf interface {
	AsIter() OfIntIterIf
}

type OfIntLooper interface {
	LoopItem(i int, d int) bool
}
type OfIntMapper interface {
	MapItem(i int, d int) int
}
type OfIntLoopFunc func(i int, d int) bool

func (__ OfIntLoopFunc) LoopItem(i int, d int) bool {
	return __(i, d)
}

type OfIntMapFunc func(i int, d int) int

func (__ OfIntMapFunc) MapItem(i int, d int) int {
	return __(i, d)
}

type OfIntIterIf interface {
	Range(fntr OfIntLooper)
	Map(fntr OfIntMapper) OfIntMutIf
}

func OfIntInto(__ OfIntIf) []int {
	switch d := __.(type) {
	case OfInt:
		return []int(d)
	case *OfIntSt:
		return []int(d._OfInt)
	case nil:
		return nil
	default:
		res := make([]int, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfInt []int

func (__ OfInt) Get(i int) int {
	return __[i]
}
func (__ OfInt) Set(i int, d int) int {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfInt) Len() int {
	return len(__)
}

func (__ OfInt) AsIter() OfIntIterIf {
	return OfIntIter(__)
}

type _OfInt = OfInt
type OfIntSt struct {
	// do not want to export but want to use embedding method
	_OfInt
}

func NewOfIntStFrom(list []int) *OfIntSt {
	return &OfIntSt{_OfInt: OfInt(list)}
}
func NewOfIntSt(i int) *OfIntSt {
	return NewOfIntStFrom(make([]int, i))
}

type OfIntIter []int

func (__ OfIntIter) Range(fntr OfIntLooper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfIntIter) Map(fntr OfIntMapper) OfIntMutIf {
	rval := make([]int, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfInt(rval)
}

type OfInt16If interface {
	Get(int) int16
	Len() int
}
type OfInt16MutIf interface {
	OfInt16If
	Set(int, int16) int16
}
type OfInt16AsIterIf interface {
	AsIter() OfInt16IterIf
}

type OfInt16Looper interface {
	LoopItem(i int, d int16) bool
}
type OfInt16Mapper interface {
	MapItem(i int, d int16) int16
}
type OfInt16LoopFunc func(i int, d int16) bool

func (__ OfInt16LoopFunc) LoopItem(i int, d int16) bool {
	return __(i, d)
}

type OfInt16MapFunc func(i int, d int16) int16

func (__ OfInt16MapFunc) MapItem(i int, d int16) int16 {
	return __(i, d)
}

type OfInt16IterIf interface {
	Range(fntr OfInt16Looper)
	Map(fntr OfInt16Mapper) OfInt16MutIf
}

func OfInt16Into(__ OfInt16If) []int16 {
	switch d := __.(type) {
	case OfInt16:
		return []int16(d)
	case *OfInt16St:
		return []int16(d._OfInt16)
	case nil:
		return nil
	default:
		res := make([]int16, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfInt16 []int16

func (__ OfInt16) Get(i int) int16 {
	return __[i]
}
func (__ OfInt16) Set(i int, d int16) int16 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfInt16) Len() int {
	return len(__)
}

func (__ OfInt16) AsIter() OfInt16IterIf {
	return OfInt16Iter(__)
}

type _OfInt16 = OfInt16
type OfInt16St struct {
	// do not want to export but want to use embedding method
	_OfInt16
}

func NewOfInt16StFrom(list []int16) *OfInt16St {
	return &OfInt16St{_OfInt16: OfInt16(list)}
}
func NewOfInt16St(i int) *OfInt16St {
	return NewOfInt16StFrom(make([]int16, i))
}

type OfInt16Iter []int16

func (__ OfInt16Iter) Range(fntr OfInt16Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfInt16Iter) Map(fntr OfInt16Mapper) OfInt16MutIf {
	rval := make([]int16, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfInt16(rval)
}

type OfInt32If interface {
	Get(int) int32
	Len() int
}
type OfInt32MutIf interface {
	OfInt32If
	Set(int, int32) int32
}
type OfInt32AsIterIf interface {
	AsIter() OfInt32IterIf
}

type OfInt32Looper interface {
	LoopItem(i int, d int32) bool
}
type OfInt32Mapper interface {
	MapItem(i int, d int32) int32
}
type OfInt32LoopFunc func(i int, d int32) bool

func (__ OfInt32LoopFunc) LoopItem(i int, d int32) bool {
	return __(i, d)
}

type OfInt32MapFunc func(i int, d int32) int32

func (__ OfInt32MapFunc) MapItem(i int, d int32) int32 {
	return __(i, d)
}

type OfInt32IterIf interface {
	Range(fntr OfInt32Looper)
	Map(fntr OfInt32Mapper) OfInt32MutIf
}

func OfInt32Into(__ OfInt32If) []int32 {
	switch d := __.(type) {
	case OfInt32:
		return []int32(d)
	case *OfInt32St:
		return []int32(d._OfInt32)
	case nil:
		return nil
	default:
		res := make([]int32, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfInt32 []int32

func (__ OfInt32) Get(i int) int32 {
	return __[i]
}
func (__ OfInt32) Set(i int, d int32) int32 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfInt32) Len() int {
	return len(__)
}

func (__ OfInt32) AsIter() OfInt32IterIf {
	return OfInt32Iter(__)
}

type _OfInt32 = OfInt32
type OfInt32St struct {
	// do not want to export but want to use embedding method
	_OfInt32
}

func NewOfInt32StFrom(list []int32) *OfInt32St {
	return &OfInt32St{_OfInt32: OfInt32(list)}
}
func NewOfInt32St(i int) *OfInt32St {
	return NewOfInt32StFrom(make([]int32, i))
}

type OfInt32Iter []int32

func (__ OfInt32Iter) Range(fntr OfInt32Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfInt32Iter) Map(fntr OfInt32Mapper) OfInt32MutIf {
	rval := make([]int32, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfInt32(rval)
}

type OfInt64If interface {
	Get(int) int64
	Len() int
}
type OfInt64MutIf interface {
	OfInt64If
	Set(int, int64) int64
}
type OfInt64AsIterIf interface {
	AsIter() OfInt64IterIf
}

type OfInt64Looper interface {
	LoopItem(i int, d int64) bool
}
type OfInt64Mapper interface {
	MapItem(i int, d int64) int64
}
type OfInt64LoopFunc func(i int, d int64) bool

func (__ OfInt64LoopFunc) LoopItem(i int, d int64) bool {
	return __(i, d)
}

type OfInt64MapFunc func(i int, d int64) int64

func (__ OfInt64MapFunc) MapItem(i int, d int64) int64 {
	return __(i, d)
}

type OfInt64IterIf interface {
	Range(fntr OfInt64Looper)
	Map(fntr OfInt64Mapper) OfInt64MutIf
}

func OfInt64Into(__ OfInt64If) []int64 {
	switch d := __.(type) {
	case OfInt64:
		return []int64(d)
	case *OfInt64St:
		return []int64(d._OfInt64)
	case nil:
		return nil
	default:
		res := make([]int64, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfInt64 []int64

func (__ OfInt64) Get(i int) int64 {
	return __[i]
}
func (__ OfInt64) Set(i int, d int64) int64 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfInt64) Len() int {
	return len(__)
}

func (__ OfInt64) AsIter() OfInt64IterIf {
	return OfInt64Iter(__)
}

type _OfInt64 = OfInt64
type OfInt64St struct {
	// do not want to export but want to use embedding method
	_OfInt64
}

func NewOfInt64StFrom(list []int64) *OfInt64St {
	return &OfInt64St{_OfInt64: OfInt64(list)}
}
func NewOfInt64St(i int) *OfInt64St {
	return NewOfInt64StFrom(make([]int64, i))
}

type OfInt64Iter []int64

func (__ OfInt64Iter) Range(fntr OfInt64Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfInt64Iter) Map(fntr OfInt64Mapper) OfInt64MutIf {
	rval := make([]int64, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfInt64(rval)
}

type OfInt8If interface {
	Get(int) int8
	Len() int
}
type OfInt8MutIf interface {
	OfInt8If
	Set(int, int8) int8
}
type OfInt8AsIterIf interface {
	AsIter() OfInt8IterIf
}

type OfInt8Looper interface {
	LoopItem(i int, d int8) bool
}
type OfInt8Mapper interface {
	MapItem(i int, d int8) int8
}
type OfInt8LoopFunc func(i int, d int8) bool

func (__ OfInt8LoopFunc) LoopItem(i int, d int8) bool {
	return __(i, d)
}

type OfInt8MapFunc func(i int, d int8) int8

func (__ OfInt8MapFunc) MapItem(i int, d int8) int8 {
	return __(i, d)
}

type OfInt8IterIf interface {
	Range(fntr OfInt8Looper)
	Map(fntr OfInt8Mapper) OfInt8MutIf
}

func OfInt8Into(__ OfInt8If) []int8 {
	switch d := __.(type) {
	case OfInt8:
		return []int8(d)
	case *OfInt8St:
		return []int8(d._OfInt8)
	case nil:
		return nil
	default:
		res := make([]int8, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfInt8 []int8

func (__ OfInt8) Get(i int) int8 {
	return __[i]
}
func (__ OfInt8) Set(i int, d int8) int8 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfInt8) Len() int {
	return len(__)
}

func (__ OfInt8) AsIter() OfInt8IterIf {
	return OfInt8Iter(__)
}

type _OfInt8 = OfInt8
type OfInt8St struct {
	// do not want to export but want to use embedding method
	_OfInt8
}

func NewOfInt8StFrom(list []int8) *OfInt8St {
	return &OfInt8St{_OfInt8: OfInt8(list)}
}
func NewOfInt8St(i int) *OfInt8St {
	return NewOfInt8StFrom(make([]int8, i))
}

type OfInt8Iter []int8

func (__ OfInt8Iter) Range(fntr OfInt8Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfInt8Iter) Map(fntr OfInt8Mapper) OfInt8MutIf {
	rval := make([]int8, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfInt8(rval)
}

type OfRuneIf interface {
	Get(int) rune
	Len() int
}
type OfRuneMutIf interface {
	OfRuneIf
	Set(int, rune) rune
}
type OfRuneAsIterIf interface {
	AsIter() OfRuneIterIf
}

type OfRuneLooper interface {
	LoopItem(i int, d rune) bool
}
type OfRuneMapper interface {
	MapItem(i int, d rune) rune
}
type OfRuneLoopFunc func(i int, d rune) bool

func (__ OfRuneLoopFunc) LoopItem(i int, d rune) bool {
	return __(i, d)
}

type OfRuneMapFunc func(i int, d rune) rune

func (__ OfRuneMapFunc) MapItem(i int, d rune) rune {
	return __(i, d)
}

type OfRuneIterIf interface {
	Range(fntr OfRuneLooper)
	Map(fntr OfRuneMapper) OfRuneMutIf
}

func OfRuneInto(__ OfRuneIf) []rune {
	switch d := __.(type) {
	case OfRune:
		return []rune(d)
	case *OfRuneSt:
		return []rune(d._OfRune)
	case nil:
		return nil
	default:
		res := make([]rune, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfRune []rune

func (__ OfRune) Get(i int) rune {
	return __[i]
}
func (__ OfRune) Set(i int, d rune) rune {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfRune) Len() int {
	return len(__)
}

func (__ OfRune) AsIter() OfRuneIterIf {
	return OfRuneIter(__)
}

type _OfRune = OfRune
type OfRuneSt struct {
	// do not want to export but want to use embedding method
	_OfRune
}

func NewOfRuneStFrom(list []rune) *OfRuneSt {
	return &OfRuneSt{_OfRune: OfRune(list)}
}
func NewOfRuneSt(i int) *OfRuneSt {
	return NewOfRuneStFrom(make([]rune, i))
}

type OfRuneIter []rune

func (__ OfRuneIter) Range(fntr OfRuneLooper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfRuneIter) Map(fntr OfRuneMapper) OfRuneMutIf {
	rval := make([]rune, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfRune(rval)
}

type OfStringIf interface {
	Get(int) string
	Len() int
}
type OfStringMutIf interface {
	OfStringIf
	Set(int, string) string
}
type OfStringAsIterIf interface {
	AsIter() OfStringIterIf
}

type OfStringLooper interface {
	LoopItem(i int, d string) bool
}
type OfStringMapper interface {
	MapItem(i int, d string) string
}
type OfStringLoopFunc func(i int, d string) bool

func (__ OfStringLoopFunc) LoopItem(i int, d string) bool {
	return __(i, d)
}

type OfStringMapFunc func(i int, d string) string

func (__ OfStringMapFunc) MapItem(i int, d string) string {
	return __(i, d)
}

type OfStringIterIf interface {
	Range(fntr OfStringLooper)
	Map(fntr OfStringMapper) OfStringMutIf
}

func OfStringInto(__ OfStringIf) []string {
	switch d := __.(type) {
	case OfString:
		return []string(d)
	case *OfStringSt:
		return []string(d._OfString)
	case nil:
		return nil
	default:
		res := make([]string, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfString []string

func (__ OfString) Get(i int) string {
	return __[i]
}
func (__ OfString) Set(i int, d string) string {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfString) Len() int {
	return len(__)
}

func (__ OfString) AsIter() OfStringIterIf {
	return OfStringIter(__)
}

type _OfString = OfString
type OfStringSt struct {
	// do not want to export but want to use embedding method
	_OfString
}

func NewOfStringStFrom(list []string) *OfStringSt {
	return &OfStringSt{_OfString: OfString(list)}
}
func NewOfStringSt(i int) *OfStringSt {
	return NewOfStringStFrom(make([]string, i))
}

type OfStringIter []string

func (__ OfStringIter) Range(fntr OfStringLooper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfStringIter) Map(fntr OfStringMapper) OfStringMutIf {
	rval := make([]string, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfString(rval)
}

type OfUintIf interface {
	Get(int) uint
	Len() int
}
type OfUintMutIf interface {
	OfUintIf
	Set(int, uint) uint
}
type OfUintAsIterIf interface {
	AsIter() OfUintIterIf
}

type OfUintLooper interface {
	LoopItem(i int, d uint) bool
}
type OfUintMapper interface {
	MapItem(i int, d uint) uint
}
type OfUintLoopFunc func(i int, d uint) bool

func (__ OfUintLoopFunc) LoopItem(i int, d uint) bool {
	return __(i, d)
}

type OfUintMapFunc func(i int, d uint) uint

func (__ OfUintMapFunc) MapItem(i int, d uint) uint {
	return __(i, d)
}

type OfUintIterIf interface {
	Range(fntr OfUintLooper)
	Map(fntr OfUintMapper) OfUintMutIf
}

func OfUintInto(__ OfUintIf) []uint {
	switch d := __.(type) {
	case OfUint:
		return []uint(d)
	case *OfUintSt:
		return []uint(d._OfUint)
	case nil:
		return nil
	default:
		res := make([]uint, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfUint []uint

func (__ OfUint) Get(i int) uint {
	return __[i]
}
func (__ OfUint) Set(i int, d uint) uint {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfUint) Len() int {
	return len(__)
}

func (__ OfUint) AsIter() OfUintIterIf {
	return OfUintIter(__)
}

type _OfUint = OfUint
type OfUintSt struct {
	// do not want to export but want to use embedding method
	_OfUint
}

func NewOfUintStFrom(list []uint) *OfUintSt {
	return &OfUintSt{_OfUint: OfUint(list)}
}
func NewOfUintSt(i int) *OfUintSt {
	return NewOfUintStFrom(make([]uint, i))
}

type OfUintIter []uint

func (__ OfUintIter) Range(fntr OfUintLooper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfUintIter) Map(fntr OfUintMapper) OfUintMutIf {
	rval := make([]uint, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfUint(rval)
}

type OfUint16If interface {
	Get(int) uint16
	Len() int
}
type OfUint16MutIf interface {
	OfUint16If
	Set(int, uint16) uint16
}
type OfUint16AsIterIf interface {
	AsIter() OfUint16IterIf
}

type OfUint16Looper interface {
	LoopItem(i int, d uint16) bool
}
type OfUint16Mapper interface {
	MapItem(i int, d uint16) uint16
}
type OfUint16LoopFunc func(i int, d uint16) bool

func (__ OfUint16LoopFunc) LoopItem(i int, d uint16) bool {
	return __(i, d)
}

type OfUint16MapFunc func(i int, d uint16) uint16

func (__ OfUint16MapFunc) MapItem(i int, d uint16) uint16 {
	return __(i, d)
}

type OfUint16IterIf interface {
	Range(fntr OfUint16Looper)
	Map(fntr OfUint16Mapper) OfUint16MutIf
}

func OfUint16Into(__ OfUint16If) []uint16 {
	switch d := __.(type) {
	case OfUint16:
		return []uint16(d)
	case *OfUint16St:
		return []uint16(d._OfUint16)
	case nil:
		return nil
	default:
		res := make([]uint16, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfUint16 []uint16

func (__ OfUint16) Get(i int) uint16 {
	return __[i]
}
func (__ OfUint16) Set(i int, d uint16) uint16 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfUint16) Len() int {
	return len(__)
}

func (__ OfUint16) AsIter() OfUint16IterIf {
	return OfUint16Iter(__)
}

type _OfUint16 = OfUint16
type OfUint16St struct {
	// do not want to export but want to use embedding method
	_OfUint16
}

func NewOfUint16StFrom(list []uint16) *OfUint16St {
	return &OfUint16St{_OfUint16: OfUint16(list)}
}
func NewOfUint16St(i int) *OfUint16St {
	return NewOfUint16StFrom(make([]uint16, i))
}

type OfUint16Iter []uint16

func (__ OfUint16Iter) Range(fntr OfUint16Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfUint16Iter) Map(fntr OfUint16Mapper) OfUint16MutIf {
	rval := make([]uint16, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfUint16(rval)
}

type OfUint32If interface {
	Get(int) uint32
	Len() int
}
type OfUint32MutIf interface {
	OfUint32If
	Set(int, uint32) uint32
}
type OfUint32AsIterIf interface {
	AsIter() OfUint32IterIf
}

type OfUint32Looper interface {
	LoopItem(i int, d uint32) bool
}
type OfUint32Mapper interface {
	MapItem(i int, d uint32) uint32
}
type OfUint32LoopFunc func(i int, d uint32) bool

func (__ OfUint32LoopFunc) LoopItem(i int, d uint32) bool {
	return __(i, d)
}

type OfUint32MapFunc func(i int, d uint32) uint32

func (__ OfUint32MapFunc) MapItem(i int, d uint32) uint32 {
	return __(i, d)
}

type OfUint32IterIf interface {
	Range(fntr OfUint32Looper)
	Map(fntr OfUint32Mapper) OfUint32MutIf
}

func OfUint32Into(__ OfUint32If) []uint32 {
	switch d := __.(type) {
	case OfUint32:
		return []uint32(d)
	case *OfUint32St:
		return []uint32(d._OfUint32)
	case nil:
		return nil
	default:
		res := make([]uint32, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfUint32 []uint32

func (__ OfUint32) Get(i int) uint32 {
	return __[i]
}
func (__ OfUint32) Set(i int, d uint32) uint32 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfUint32) Len() int {
	return len(__)
}

func (__ OfUint32) AsIter() OfUint32IterIf {
	return OfUint32Iter(__)
}

type _OfUint32 = OfUint32
type OfUint32St struct {
	// do not want to export but want to use embedding method
	_OfUint32
}

func NewOfUint32StFrom(list []uint32) *OfUint32St {
	return &OfUint32St{_OfUint32: OfUint32(list)}
}
func NewOfUint32St(i int) *OfUint32St {
	return NewOfUint32StFrom(make([]uint32, i))
}

type OfUint32Iter []uint32

func (__ OfUint32Iter) Range(fntr OfUint32Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfUint32Iter) Map(fntr OfUint32Mapper) OfUint32MutIf {
	rval := make([]uint32, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfUint32(rval)
}

type OfUint64If interface {
	Get(int) uint64
	Len() int
}
type OfUint64MutIf interface {
	OfUint64If
	Set(int, uint64) uint64
}
type OfUint64AsIterIf interface {
	AsIter() OfUint64IterIf
}

type OfUint64Looper interface {
	LoopItem(i int, d uint64) bool
}
type OfUint64Mapper interface {
	MapItem(i int, d uint64) uint64
}
type OfUint64LoopFunc func(i int, d uint64) bool

func (__ OfUint64LoopFunc) LoopItem(i int, d uint64) bool {
	return __(i, d)
}

type OfUint64MapFunc func(i int, d uint64) uint64

func (__ OfUint64MapFunc) MapItem(i int, d uint64) uint64 {
	return __(i, d)
}

type OfUint64IterIf interface {
	Range(fntr OfUint64Looper)
	Map(fntr OfUint64Mapper) OfUint64MutIf
}

func OfUint64Into(__ OfUint64If) []uint64 {
	switch d := __.(type) {
	case OfUint64:
		return []uint64(d)
	case *OfUint64St:
		return []uint64(d._OfUint64)
	case nil:
		return nil
	default:
		res := make([]uint64, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfUint64 []uint64

func (__ OfUint64) Get(i int) uint64 {
	return __[i]
}
func (__ OfUint64) Set(i int, d uint64) uint64 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfUint64) Len() int {
	return len(__)
}

func (__ OfUint64) AsIter() OfUint64IterIf {
	return OfUint64Iter(__)
}

type _OfUint64 = OfUint64
type OfUint64St struct {
	// do not want to export but want to use embedding method
	_OfUint64
}

func NewOfUint64StFrom(list []uint64) *OfUint64St {
	return &OfUint64St{_OfUint64: OfUint64(list)}
}
func NewOfUint64St(i int) *OfUint64St {
	return NewOfUint64StFrom(make([]uint64, i))
}

type OfUint64Iter []uint64

func (__ OfUint64Iter) Range(fntr OfUint64Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfUint64Iter) Map(fntr OfUint64Mapper) OfUint64MutIf {
	rval := make([]uint64, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfUint64(rval)
}

type OfUint8If interface {
	Get(int) uint8
	Len() int
}
type OfUint8MutIf interface {
	OfUint8If
	Set(int, uint8) uint8
}
type OfUint8AsIterIf interface {
	AsIter() OfUint8IterIf
}

type OfUint8Looper interface {
	LoopItem(i int, d uint8) bool
}
type OfUint8Mapper interface {
	MapItem(i int, d uint8) uint8
}
type OfUint8LoopFunc func(i int, d uint8) bool

func (__ OfUint8LoopFunc) LoopItem(i int, d uint8) bool {
	return __(i, d)
}

type OfUint8MapFunc func(i int, d uint8) uint8

func (__ OfUint8MapFunc) MapItem(i int, d uint8) uint8 {
	return __(i, d)
}

type OfUint8IterIf interface {
	Range(fntr OfUint8Looper)
	Map(fntr OfUint8Mapper) OfUint8MutIf
}

func OfUint8Into(__ OfUint8If) []uint8 {
	switch d := __.(type) {
	case OfUint8:
		return []uint8(d)
	case *OfUint8St:
		return []uint8(d._OfUint8)
	case nil:
		return nil
	default:
		res := make([]uint8, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfUint8 []uint8

func (__ OfUint8) Get(i int) uint8 {
	return __[i]
}
func (__ OfUint8) Set(i int, d uint8) uint8 {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfUint8) Len() int {
	return len(__)
}

func (__ OfUint8) AsIter() OfUint8IterIf {
	return OfUint8Iter(__)
}

type _OfUint8 = OfUint8
type OfUint8St struct {
	// do not want to export but want to use embedding method
	_OfUint8
}

func NewOfUint8StFrom(list []uint8) *OfUint8St {
	return &OfUint8St{_OfUint8: OfUint8(list)}
}
func NewOfUint8St(i int) *OfUint8St {
	return NewOfUint8StFrom(make([]uint8, i))
}

type OfUint8Iter []uint8

func (__ OfUint8Iter) Range(fntr OfUint8Looper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfUint8Iter) Map(fntr OfUint8Mapper) OfUint8MutIf {
	rval := make([]uint8, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfUint8(rval)
}

type OfUintptrIf interface {
	Get(int) uintptr
	Len() int
}
type OfUintptrMutIf interface {
	OfUintptrIf
	Set(int, uintptr) uintptr
}
type OfUintptrAsIterIf interface {
	AsIter() OfUintptrIterIf
}

type OfUintptrLooper interface {
	LoopItem(i int, d uintptr) bool
}
type OfUintptrMapper interface {
	MapItem(i int, d uintptr) uintptr
}
type OfUintptrLoopFunc func(i int, d uintptr) bool

func (__ OfUintptrLoopFunc) LoopItem(i int, d uintptr) bool {
	return __(i, d)
}

type OfUintptrMapFunc func(i int, d uintptr) uintptr

func (__ OfUintptrMapFunc) MapItem(i int, d uintptr) uintptr {
	return __(i, d)
}

type OfUintptrIterIf interface {
	Range(fntr OfUintptrLooper)
	Map(fntr OfUintptrMapper) OfUintptrMutIf
}

func OfUintptrInto(__ OfUintptrIf) []uintptr {
	switch d := __.(type) {
	case OfUintptr:
		return []uintptr(d)
	case *OfUintptrSt:
		return []uintptr(d._OfUintptr)
	case nil:
		return nil
	default:
		res := make([]uintptr, __.Len())
		for i := 0; i < len(res); i += 1 {
			res[i] = __.Get(i)
		}
		return res
	}
}

type OfUintptr []uintptr

func (__ OfUintptr) Get(i int) uintptr {
	return __[i]
}
func (__ OfUintptr) Set(i int, d uintptr) uintptr {
	old := __[i]
	__[i] = d
	return old
}

func (__ OfUintptr) Len() int {
	return len(__)
}

func (__ OfUintptr) AsIter() OfUintptrIterIf {
	return OfUintptrIter(__)
}

type _OfUintptr = OfUintptr
type OfUintptrSt struct {
	// do not want to export but want to use embedding method
	_OfUintptr
}

func NewOfUintptrStFrom(list []uintptr) *OfUintptrSt {
	return &OfUintptrSt{_OfUintptr: OfUintptr(list)}
}
func NewOfUintptrSt(i int) *OfUintptrSt {
	return NewOfUintptrStFrom(make([]uintptr, i))
}

type OfUintptrIter []uintptr

func (__ OfUintptrIter) Range(fntr OfUintptrLooper) {
	for i := range __ {
		if !fntr.LoopItem(i, __[i]) {
			break
		}
	}
}
func (__ OfUintptrIter) Map(fntr OfUintptrMapper) OfUintptrMutIf {
	rval := make([]uintptr, len(__))
	for i := range __ {
		rval[i] = fntr.MapItem(i, __[i])
	}
	return OfUintptr(rval)
}
