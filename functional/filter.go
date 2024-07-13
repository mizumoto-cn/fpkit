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

// Filter returns a new slice containing only the elements that satisfy the predicate.
//	Filter(func(x int, i int) bool { return x > 0 }, 1, -2, 3, -4, 5) // [1, 3, 5]
func Filter[T any](fn func(T, int) bool, input ...T) []T {
	list := make([]T, len(input))
	newLen := 0
	for i := range input {
		if fn(input[i], i) {
			newLen++
			list[newLen-1] = input[i]
		}
	}
	result := list[:newLen]
	return result
}
