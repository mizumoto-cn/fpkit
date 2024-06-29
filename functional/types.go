package functional

// basic type for function
type FnObject func(any) any

// return true if src is "defined to be equal" to dst
type eqFn[T any] func(src, dst T) bool

// return true if src is "defined to match" a given value or condition
type matchFn[T any] func(src T) bool

var _ = FnObject(nil)
var _ = eqFn[any](nil)
var _ = matchFn[any](nil)
