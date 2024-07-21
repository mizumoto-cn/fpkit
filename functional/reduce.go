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

// Foldl applies a function to each element of a list, starting from the left, and returns the final value.
//	index := 0
//	Foldl([]int{1, 2, 3, 4, 5, 6}, func(acc, x int) int {
//		index++
//		if index % 2 == 0 {
//			return acc - x
//		}
//		return acc + x
//	}, 0) // 0 + 1 - 2 + 3 - 4 + 5 - 6 = -3
func Foldl[T any, U any](s []T, fn func(U, T) U, init U) U {
	for _, v := range s {
		init = fn(init, v)
	}
	return init
}

// Foldr applies a function to each element of a list, starting from the right, and returns the final value.
//	index := 0
//	Foldr([]int{1, 2, 3, 4, 5, 6}, func(acc, x int) int {
//		index++
//		if index % 2 == 0 {
//			return acc - x
//		}
//		return acc + x
//	}, 0) // 0 + 6 - 5 + 4 - 3 + 2 - 1 = 3
func Foldr[T any, U any](s []T, fn func(U, T) U, init U) U {
	for i := len(s) - 1; i >= 0; i-- {
		init = fn(init, s[i])
	}
	return init
}

// Reduce applies a reduction function to each element of the slice on a left-to-right basis to a given initial value.
//	Reduce([]int{1, 2, 3, 4, 5, 6}, func(acc, x int) int {
//		return acc + x
//	}, 0) // 0 + 1 + 2 + 3 + 4 + 5 + 6 = 21
func Reduce[T any, U any](s []T, fn func(U, T) U, init U) U {
	// result := init
	// for _, v := range s {
	// 	result = fn(result, v)
	// }
	// return result
	return Foldl(s, fn, init)
}
