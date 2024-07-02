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

func Sum[T Numeric](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

// For comparison in this package, >0 means a > b, <0 means a < b, and 0 means a == b.
func CompareTo[T Orderable](a, b T) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func Sort[T Orderable](a []T, cmp Comparator[T]) {
	sort.SliceStable(a, func(i, j int) bool {
		return cmp(a[i], a[j])
	})
}

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

func SortAsc[T Orderable](input ...T) []T {
	return SortOrdered(true, input...)
}

func SortDesc[T Orderable](input ...T) []T {
	return SortOrdered(false, input...)
}
