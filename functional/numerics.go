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

import "sort"

// Sum returns the sum of the given numbers.
//	Sum(1, 2, 3) // => 6
func Sum[T Numeric](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

// For comparison in this package, >0 means a > b, <0 means a < b, and 0 means a == b.
//	CompareTo(1, 2) // => -1
func CompareTo[T Orderable](a, b T) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// Sort sorts the given slice in place using the given comparator.
//	Sort([]int{3, 1, 2}, func(a, b int) bool { return a < b }) // => []int{1, 2, 3}
func Sort[T Orderable](a []T, cmp Comparator[T]) {
	sort.SliceStable(a, func(i, j int) bool {
		return cmp(a[i], a[j])
	})
}

// SortOrdered sorts the given slice in ascending or descending order.
//	SortOrdered(true, 3, 1, 2) // => []int{1, 2, 3}
func SortOrdered[T Orderable](ascending bool, input ...T) []T {
	if ascending {
		Sort(input, func(a, b T) bool {
			return CompareTo(a, b) < 0
			// a < b, ascending
		})
	} else {
		Sort(input, func(a, b T) bool {
			return CompareTo(a, b) > 0
			// a > b, descending
		})
	}
	return input
}

// SortAsc sorts the given slice in ascending order.
//	SortAsc(3, 1, 2) // => []int{1, 2, 3}
func SortAsc[T Orderable](input ...T) []T {
	return SortOrdered(true, input...)
}

// SortDesc sorts the given slice in descending order.
//	SortDesc(3, 1, 2) // => []int{3, 2, 1}
func SortDesc[T Orderable](input ...T) []T {
	return SortOrdered(false, input...)
}
