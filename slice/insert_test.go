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

	"github.com/mizumoto-cn/fpkit/internal/err"
	"github.com/mizumoto-cn/fpkit/slice"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		title       string
		slice       []int
		index       int
		value       int
		want        []int
		expectedErr error
	}{
		{
			title: "insert to the head",
			slice: []int{1, 2, 3},
			index: 0,
			value: 0,
			want:  []int{0, 1, 2, 3},
		},
		{
			title: "insert to the middle",
			slice: []int{1, 2, 3},
			index: 1,
			value: 4,
			want:  []int{1, 4, 2, 3},
		},
		{
			title: "insert to the tail",
			slice: []int{1, 2, 3},
			index: 3,
			value: 4,
			want:  []int{1, 2, 3, 4},
		},
		{
			title:       "insert to the out of range",
			slice:       []int{1, 2, 3},
			index:       4,
			value:       4,
			expectedErr: err.NewIndexOutOfRangeError(4, 3),
		},
		{
			title:       "insert to the negative index",
			slice:       []int{1, 2, 3},
			index:       -1,
			value:       4,
			expectedErr: err.NewIndexOutOfRangeError(-1, 3),
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got, err := slice.Insert(c.slice, c.index, c.value)
			if c.expectedErr != nil {
				assert.Equal(t, c.expectedErr, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}
