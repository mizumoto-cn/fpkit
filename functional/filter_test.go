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
package functional_test

import (
	"testing"

	"github.com/mizumoto-cn/fpkit/functional"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		title string
		src   []int
		fn    func(int, int) bool
		want  []int
	}{
		{
			title: "empty",
			src:   []int{},
			fn: func(i int, _ int) bool {
				return i%2 == 0
			},
			want: []int{},
		},
		{
			title: "single",
			src:   []int{1},
			fn: func(i int, _ int) bool {
				return i%2 == 0
			},
			want: []int{},
		},
		{
			title: "multiple",
			src:   []int{1, 2, 3, 4, 5},
			fn: func(i int, _ int) bool {
				return i%2 == 0
			},
			want: []int{2, 4},
		},
		{
			title: "index",
			src:   []int{1, 2, 3, 4, 5},
			fn: func(i int, j int) bool {
				return i%2 == 1 && j%2 == 0
			},
			want: []int{1, 3, 5},
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got := functional.Filter(c.fn, c.src...)
			assert.Equal(t, c.want, got)
		})
	}
}
