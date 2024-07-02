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

func TestDelete(t *testing.T) {
	cases := []struct {
		name  string
		s     []int
		index int
		want  []int
		err   bool
	}{
		{
			name:  "Delete an element from the slice",
			s:     []int{1, 2, 3, 4, 5},
			index: 2,
			want:  []int{1, 2, 4, 5},
			err:   false,
		},
		{
			name:  "Delete an element from the slice with the first index",
			s:     []int{1, 2, 3, 4, 5},
			index: 0,
			want:  []int{2, 3, 4, 5},
			err:   false,
		},
		{
			name:  "Delete an element from the slice with the last index",
			s:     []int{1, 2, 3, 4, 5},
			index: 4,
			want:  []int{1, 2, 3, 4},
			err:   false,
		},
		{
			name:  "Delete an element from the slice with an invalid index",
			s:     []int{1, 2, 3, 4, 5},
			index: 5,
			want:  nil,
			err:   true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := slice.Delete(tc.s, tc.index)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestDeleteMatched(t *testing.T) {
	match := func(v int) bool {
		return v%2 == 0
	}
	cases := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "Delete all even numbers from the slice 1",
			s:    []int{1, 2, 3, 4, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "Delete all even numbers from the slice 2",
			s:    []int{2, 4, 6, 8, 10},
			want: []int{},
		},
		{
			name: "Delete all even numbers from the slice 3",
			s:    []int{1, 3, 5, 7, 9},
			want: []int{1, 3, 5, 7, 9},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := slice.DeleteMatched(tc.s, match)
			assert.Equal(t, tc.want, got)
		})
	}
}
