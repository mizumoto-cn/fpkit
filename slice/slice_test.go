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

func TestToSlice(t *testing.T) {
	cases := []struct {
		title string
		v     []any
		want  []any
	}{
		{
			title: "ints",
			v:     []any{1, 2, 3},
			want:  []any{1, 2, 3},
		},
		{
			title: "strings",
			v:     []any{"a", "b", "c"},
			want:  []any{"a", "b", "c"},
		},
		{
			title: "structs",
			v:     []any{struct{ Name string }{"mizumoto"}, struct{ Name string }{"cn"}},
			want:  []any{struct{ Name string }{"mizumoto"}, struct{ Name string }{"cn"}},
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got := slice.ToSlice(c.v...)
			assert.Equal(t, c.want, got)
		})
	}
}
