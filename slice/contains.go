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

import "github.com/mizumoto-cn/fpkit/functional"

// Contains returns true if the slice contains the value.
func Contains[T comparable](s []T, v T) bool {
	return ContainsFunc(s, func(x T) bool { return x == v })
}

// ContainsFunc returns true if the slice contains an element that satisfies the predicate.
func ContainsFunc[T any](s []T, eq func(T) bool) bool {
	for _, v := range s {
		if eq(v) {
			return true
		}
	}
	return false
}

// ContainsAny returns true if the slice contains any of the values.
func ContainsAny[T comparable](s []T, vs ...T) bool {
	for _, v := range vs {
		if Contains(s, v) {
			return true
		}
	}
	return false
}

// ContainsAll checks if all elements of subset are in the main slice.
func ContainsAll[T comparable](mainSlice, subset []T) bool {
	// Use a map to track occurrences in the main slice for quick lookup
	elementCount := make(map[T]int)
	for _, elem := range mainSlice {
		elementCount[elem]++
	}

	// Use Reduce to check if all elements of the subset are in the main slice
	return functional.Reduce(subset, func(acc bool, elem T) bool {
		if acc && elementCount[elem] > 0 {
			elementCount[elem]--
			return true
		}
		return false
	}, true)
}
