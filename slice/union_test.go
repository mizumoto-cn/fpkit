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
package slice_test

import (
	"testing"

	"github.com/mizumoto-cn/fpkit/functional"
	"github.com/mizumoto-cn/fpkit/slice"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	cases := []struct {
		name string
		s1   []int
		s2   []int
		want []int
	}{
		{
			name: "Union two slices",
			s1:   []int{1, 2, 3},
			s2:   []int{3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Union two slices with the same elements",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Union two slices with the same elements in different orders",
			s1:   []int{1, 2, 3},
			s2:   []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "Union two empty slices",
			s1:   []int{},
			s2:   []int{},
			want: []int{},
		},
		{
			name: "Union an empty slice and a non-empty slice",
			s1:   []int{},
			s2:   []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Union a non-empty slice and an empty slice",
			s1:   []int{1, 2, 3},
			s2:   []int{},
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := slice.Union(tc.s1, tc.s2)
			// The order of the elements in the result is not guaranteed
			got = functional.SortAsc(got...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		name string
		s1   []int
		s2   []int
		want []int
	}{
		{
			name: "Intersection two slices",
			s1:   []int{1, 2, 3},
			s2:   []int{3, 4, 5},
			want: []int{3},
		},
		{
			name: "Intersection two slices with the same elements",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Intersection two slices with the same elements in different orders",
			s1:   []int{1, 2, 3},
			s2:   []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "Intersection two empty slices",
			s1:   []int{},
			s2:   []int{},
			want: []int{},
		},
		{
			name: "Intersection an empty slice and a non-empty slice",
			s1:   []int{},
			s2:   []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "Intersection a non-empty slice and an empty slice",
			s1:   []int{1, 2, 3},
			s2:   []int{},
			want: []int{},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := slice.Intersection(tc.s1, tc.s2)
			// The order of the elements in the result is not guaranteed
			got = functional.SortAsc(got...)
			assert.Equal(t, tc.want, got)
		})
	}
}
