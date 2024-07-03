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

	IsTypeOf(reflect.Type) bool
	IsKindOf(reflect.Kind) bool

	Unwrap() T
	UnwrapAny() any
}

type maybe[T any] struct {
	value T
	isNil bool
}

type none[T any] struct {
	maybe[T]
}

var Maybe Optional[any] = maybe[any]{}

var None Optional[any] = none[any]{IsNil: true, value: nil}

func Just[T any](value T) Optional[T] {
	isNil := IsNil(value)
	return maybe[T]{value: value, isNil: isNil}
}

// MakeClone: make a clone of the Optional object
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
func (m maybe[T]) Just(value any) Optional[any] {
	if IsNil(value) {
		return None
	}
	return Just(value)
}

// OrElse: return the value if present, otherwise return the default value of Type T
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

// FlatMap
