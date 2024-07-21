/*
 * Copyright (c) 2024 Ruiyuan "mizumoto-cn" Xu
 *
 * This file is part of "github.com/mizumoto-cn/fpkit".
 *
 * Licensed under the Mizumoto General Public License v1.5 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://github.com/mizumoto-cn/fpkit/blob/main/LICENSE
 *     https://github.com/mizumoto-cn/fpkit/blob/main/licensing
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package functional

import "reflect"

type Optional[T any] interface {
	// True if the value is present, not nil
	IsPresent() bool
	// True if the value is not present, nil
	IsNil() bool
	// True if the value is valid
	IsValid() bool
	// True if the value is a pointer
	IsPtr() bool

	Just(any) Optional[any]
	// Just(T) Optional[T]
	// Just(any) Optional[any] to match None object

	// OrElse: return the value if present, otherwise return the default value of Type T
	OrElse(T) T

	Clone() Optional[T]

	// FlatMap
	//
	// --  >>= :: Maybe a -> (a -> Maybe b) -> Maybe b
	// exampleFlatMap :: Maybe Double -> Maybe Double
	// exampleFlatMap maybeValue = maybeValue >>= (\x -> safeDivide x 2)
	//
	// -- Example of using flatMap
	// main :: IO ()
	// main = do
	//     let result1 = safeDivide 10 2 >>= (\x -> safeDivide x 2)
	//     let result2 = safeDivide 10 0 >>= (\x -> safeDivide x 2)
	//     putStrLn (exampleMaybe result1)  -- "Result is 2.5"
	//     putStrLn (exampleMaybe result2)  -- "Division by zero!"
	FlatMap(func(T) Optional[T]) Optional[T]

	// IfPresent: if the value is present, then apply the function, otherwise do nothing
	IfPresent(func())

	// Kind: reflect.Kind of the value, which is the underlying type of the value, i.e. int, slice, struct, ptr, etc.
	Kind() reflect.Kind
	// Type: reflect.Type of the value, which is the detailed type of the value, i.e. int, []int, main.MyStruct, *int, etc.
	Type() reflect.Type

	IsKindOf(reflect.Kind) bool
	IsTypeOf(reflect.Type) bool

	Unwrap() T
	UnwrapAny() any
}

type maybe[T any] struct {
	value T
	isNil bool
}

type none struct {
	maybe[any]
}

// Maybe: maybe, Optional[any]
//	j := Maybe.Just(42) // j is a Optional[any] object with value 42
//	k := j.OrElse(0)    // k is 42
//	k = j.Unwrap()      // k is 42
var Maybe Optional[any] = maybe[any]{}

// None: none, Optional[any]
//	j := Maybe.Just(nil) // j is a Optional[any], also a none.
//	k := j.OrElse(0)    // k is 0
//	l := Just(nil)      // l is a Optional[any] object with value nil, also a none.
//	k = l.UnwrapAny()      // k is nil
var None Optional[any] = none{maybe[any]{isNil: true, value: nil}}

// Just: Just(T) Optional[T]
//	j := Just(42) // j is a Optional[int] object with value 42
//	k := j.OrElse(0)    // k is 42
//	k = j.Unwrap()      // k is 42
func Just[T any](value T) Optional[T] {
	isNil := IsNil(value)
	return maybe[T]{value: value, isNil: isNil}
}

// MakeClone: make a clone of the Optional object
//	j := Just(42) // j is a Optional[int] object with value 42
//	ptr := new(int)
//	k := MakeClone(j, ptr) // k is a Optional[int] object with value 42, while value 42 stored in the address ptr points to.
func MakeClone[T any](m Optional[T], dest *T) Optional[T] {
	if m.IsNil() {
		// return a new Optional object with nil value if the original Optional object is nil
		return Just(m.Unwrap())
	}

	// use reflection to clone the value
	x := reflect.ValueOf(m.Unwrap())
	if x.Kind() == reflect.Ptr {
		// if the value is a pointer, then clone the value by creating a new pointer
		starX := x.Elem()
		y := reflect.New(starX.Type())
		starY := y.Elem()
		starY.Set(starX)
		reflect.ValueOf(dest).Elem().Set(y.Elem())
		return Just(*dest)
	}

	// Otherwise, clone the value by copying the value
	*dest = x.Interface().(T)
	return Just(*dest)
}

// IsPresent: True if the value is present, not nil
func (m maybe[T]) IsPresent() bool {
	return !m.isNil
}

// IsNil: True if the value is not present, nil
func (m maybe[T]) IsNil() bool {
	return m.isNil
}

// IsValid: True if the value is valid
func (m maybe[T]) IsValid() bool {
	return reflect.ValueOf(m.value).IsValid()
}

// IsPtr: True if the value is a pointer
func (m maybe[T]) IsPtr() bool {
	return IsPtr(m.value)
}

// Just: Just(any) Optional[any]
//	j := Maybe.Just(42) // j is a Optional[any] object with value 42
//	k := j.OrElse(0)    // k is 42
//	k = j.Unwrap()      // k is 42
func (m maybe[T]) Just(value any) Optional[any] {
	if IsNil(value) {
		return None
	}
	return Just(value)
}

// OrElse: return the value if present, otherwise return the default value of Type T
//	j := Maybe.Just(42) // j is a Optional[any] object with value 42
//	k := j.OrElse(0)    // k is 42
//	k = j.Unwrap()      // k is 42
//	j = Maybe.Just(nil) // j is a Optional[any] object with value nil
//	k = j.OrElse(0)     // k is 0
//	k = j.Unwrap()      // k is nil
func (m maybe[T]) OrElse(defaultValue T) T {
	if m.IsNil() {
		return defaultValue
	}
	return m.value
}

// Clone: make a clone of the Optional object
func (m maybe[T]) Clone() Optional[T] {
	return MakeClone(m, new(T))
}

// FlatMap is a monadic operation that applies a function to the value
// and returns a new Optional object with the new value.
//	j := Maybe.Just(10) // j is a Optional[int] object with value 10
//	result := j.FlatMap(func(v int) functional.Optional[int] { return Just(v / 2) })
//	// result is a Optional[int] object with value 5
//	k := None.FlatMap(func(v any) functional.Optional[any] { return Just(42) })
//	// k is a Optional[any] object with value 42
//	l := k.Unwrap() // l is 42
func (m maybe[T]) FlatMap(fn func(T) Optional[T]) Optional[T] {
	return fn(m.value)
}

// IfPresent: if the value is present, then apply the function, otherwise do nothing
func (m maybe[T]) IfPresent(fn func()) {
	if m.IsPresent() {
		fn()
	}
}

// Kind: reflect.Kind of the value, which is the underlying type of the value, i.e. int, slice, struct, ptr, etc.
func (m maybe[T]) Kind() reflect.Kind {
	return reflect.ValueOf(m.value).Kind()
}

// Type: reflect.Type of the value, which is the detailed type of the value, i.e. int, []int, main.MyStruct, *int, etc.
func (m maybe[T]) Type() reflect.Type {
	if m.IsNil() {
		return reflect.TypeOf(nil)
	}
	return reflect.TypeOf(m.value)
}

// IsKindOf: True if the value is of the specified kind
func (m maybe[T]) IsKindOf(k reflect.Kind) bool {
	return m.Kind() == k
}

// IsTypeOf: True if the value is of the specified type
func (m maybe[T]) IsTypeOf(t reflect.Type) bool {
	return m.Type() == t
}

// Unwrap: return the value
func (m maybe[T]) Unwrap() T {
	return m.value
}

// UnwrapAny: return the value as any
func (m maybe[T]) UnwrapAny() any {
	if m.IsNil() {
		return nil
	}
	return m.value
}

// IsPresent: True if the value is present, not nil
func (n none) IsPresent() bool {
	return false
}

// IsNil: True if the value is not present, nil
func (n none) IsNil() bool {
	return true
}

// IsValid: True if the value is valid
func (n none) IsValid() bool {
	return false
}

// IsPtr: True if the value is a pointer
func (n none) IsPtr() bool {
	return false
}

// OrElse: return the value if present, otherwise return the default value of Type T
func (n none) OrElse(defaultValue any) any {
	return defaultValue
}

// Clone: make a clone of the Optional object
func (n none) Clone() Optional[any] {
	return None
}

// IfPresent: if the value is present, then apply the function, otherwise do nothing
func (n none) IfPresent(fn func()) {
}

// Kind: reflect.Kind of the value, which is the underlying type of the value, i.e. int, slice, struct, ptr, etc.
func (n none) Kind() reflect.Kind {
	return reflect.Invalid
}

// Type: reflect.Type of the value, which is the detailed type of the value, i.e. int, []int, main.MyStruct, *int, etc.
func (n none) Type() reflect.Type {
	return reflect.TypeOf(nil)
}

// IsKindOf: True if the value is of the specified kind
func (n none) IsKindOf(k reflect.Kind) bool {
	return false
}

// Unwrap: return the value
func (n none) Unwrap() any {
	return nil
}

// UnwrapAny: return the value as any
func (n none) UnwrapAny() any {
	return nil
}
