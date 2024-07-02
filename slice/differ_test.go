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

	"github.com/mizumoto-cn/fpkit/slice"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	var empty []int
	cases := []struct {
		title string
		s1    []int
		s2    []int
		want  []int
	}{
		{
			title: "Difference between two slices: partial match",
			s1:    []int{1, 2, 3, 4},
			s2:    []int{2, 3, 5, 6},
			want:  []int{1, 4},
		},
		{
			title: "Difference between two slices: full match",
			s1:    []int{1, 2, 3, 4},
			s2:    []int{1, 2, 3, 4},
			want:  empty,
		},
		{
			title: "Difference between two slices: no match",
			s1:    []int{1, 2, 3, 4},
			s2:    []int{5, 6, 7, 8},
			want:  []int{1, 2, 3, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got := slice.Difference(c.s1, c.s2)
			assert.Equal(t, c.want, got)
		})
	}
}
