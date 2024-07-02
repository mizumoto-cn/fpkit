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
package slice

// Index returns the index of the first occurrence of the value v in the slice src, or -1 if not present.
func Index[T comparable](src []T, v T) int {
	return IndexMatchFunc(src, func(e T) bool { return e == v })
}

// IndexMatchFunc returns the index of the first element in the slice that satisfies the provided function, or -1 if not present.
func IndexMatchFunc[T any](src []T, match func(T) bool) int {
	for k, v := range src {
		if match(v) {
			return k
		}
	}
	return -1
}

// IndexAllMatchFunc returns the index of the all elements in the slice that satisfies the provided function.
func IndexAllMatchFunc[T any](src []T, match func(T) bool) []int {
	var indexes []int
	for k, v := range src {
		if match(v) {
			indexes = append(indexes, k)
		}
	}
	return indexes
}

// IndexAll returns the index of the all elements in the slice that satisfies the provided function.
func IndexAll[T comparable](src []T, v T) []int {
	return IndexAllMatchFunc(src, func(e T) bool { return e == v })
}

// LastIndexMatchFunc returns the index of the last element in the slice that satisfies the provided function, or -1 if not present.
func LastIndexMatchFunc[T any](src []T, match func(T) bool) int {
	for i := len(src) - 1; i >= 0; i-- {
		if match(src[i]) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last occurrence of the value v in the slice src, or -1 if not present.
func LastIndex[T comparable](src []T, v T) int {
	return LastIndexMatchFunc(src, func(e T) bool { return e == v })
}
