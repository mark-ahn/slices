package slices

// type If[Some any] interface {
// 	Get(int) Some
// 	Len() int
// }
// type MutIf[Some any] interface {
// 	If[Some]
// 	Set(int, Some) Some
// }
// type AsIterIf[Some any] interface {
// 	AsIter() IterIf[Some]
// }

// type Looper[Some any] interface {
// 	LoopItem(i int, d Some) bool
// }
// type Mapper[Some any] interface {
// 	MapItem(i int, d Some) Some
// }
// type LoopFunc[Some any] func(i int, d Some) bool

// func (__ LoopFunc[Some]) LoopItem(i int, d Some) bool {
// 	return __(i, d)
// }

// type MapFunc[Some any] func(i int, d Some) Some

// func (__ MapFunc[Some]) MapItem(i int, d Some) Some {
// 	return __(i, d)
// }

// type IterIf[Some any] interface {
// 	Range(fntr Looper[Some])
// 	Map(fntr Mapper[Some]) MutIf[Some]
// }

// func Into[Some any](__ If[Some]) []Some {
// 	switch d := __.(type) {
// 	case Of[Some]:
// 		return []Some(d)
// 	case *OfSt[Some]:
// 		return []Some(d.Of)
// 	case nil:
// 		return nil
// 	default:
// 		res := make([]Some, __.Len())
// 		for i := 0; i < len(res); i += 1 {
// 			res[i] = __.Get(i)
// 		}
// 		return res
// 	}
// }

// type Of[Some any] []Some

// func (__ Of[Some]) Get(i int) Some {
// 	return __[i]
// }
// func (__ Of[Some]) Set(i int, d Some) Some {
// 	old := __[i]
// 	__[i] = d
// 	return old
// }

// func (__ Of[Some]) Len() int {
// 	return len(__)
// }

// func (__ Of[Some]) AsIter() IterIf[Some] {
// 	return Iter[Some](__)
// }

// type OfSt[Some any] struct {
// 	// do not want to export but want to use embedding method
// 	Of[Some]
// }

// func NewSliceOfStFrom[Some any](list []Some) *OfSt[Some] {
// 	return &OfSt[Some]{Of: (Of[Some](list))}
// }
// func NewSliceOfSt[Some any](i int) *OfSt[Some] {
// 	return NewSliceOfStFrom(make([]Some, i))
// }

// type Iter[Some any] []Some

// func (__ Iter[Some]) Range(fntr Looper[Some]) {
// 	for i := range __ {
// 		if !fntr.LoopItem(i, __[i]) {
// 			break
// 		}
// 	}
// }

// func (__ Iter[Some]) Map(fntr Mapper[Some]) MutIf[Some] {
// 	rval := make([]Some, len(__))
// 	for i := range __ {
// 		rval[i] = fntr.MapItem(i, __[i])
// 	}
// 	return Of[Some](rval)
// }
