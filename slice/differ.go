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

// Difference returns a slice of elements that are in s1 but not in s2.
// returns an nil slice if s1 and s2 are equal.
func Difference[T comparable](s1, s2 []T) []T {
	elementMap := make(map[T]bool)
	for _, v := range s2 {
		elementMap[v] = true
	}
	var diff []T
	for _, v := range s1 {
		if !elementMap[v] {
			diff = append(diff, v)
		}
	}
	return diff
}
