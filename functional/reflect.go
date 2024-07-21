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

// The PtrOf function returns a new pointer to a copy of the passed-in value,
// not the address of the original variable.
func PtrOf[T any](v T) *T {
	return &v
}

// The SliceOf function returns a new slice containing the passed-in values.
//	SliceOf(1, 2, 3, 4, 5) // [1, 2, 3, 4, 5]
func SliceOf[T any](v ...T) []T {
	return v
}

// The IsPtr function returns true if the passed-in value is a pointer.
//	IsPtr(42) // false
func IsPtr[T any](v T) bool {
	return Kind(v) == reflect.Ptr
}

// The Kind function returns the reflect.Kind of the passed-in value.
//	Kind(42) // int
func Kind[T any](v T) reflect.Kind {
	return reflect.ValueOf(v).Kind()
}

// The IsNil function returns true if the passed-in value is nil.
//	IsNil(nil) // true
func IsNil[T any](v T) bool {
	// 2 cases, pointer or not
	val := reflect.ValueOf(v)
	if Kind(v) == reflect.Ptr {
		return val.IsNil()
	}
	return !val.IsValid()
}
