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

func TestMax(t *testing.T) {
	cases := []struct {
		title  string
		src    []int
		want   int
		panics bool
	}{
		{
			title:  "empty",
			src:    []int{},
			want:   0,
			panics: true,
		},
		{
			title:  "single",
			src:    []int{1},
			want:   1,
			panics: false,
		},
		{
			title:  "multiple",
			src:    []int{1, 2, 3, 4, 5},
			want:   5,
			panics: false,
		},
		{
			title:  "negative",
			src:    []int{-1, -2, -3, -4, -5},
			want:   -1,
			panics: false,
		},
		{
			title:  "mixed",
			src:    []int{-1, 2, -3, 4, -5},
			want:   4,
			panics: false,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if c.panics {
				assert.Panics(t, func() { slice.Max(c.src) })
			} else {
				assert.Equal(t, c.want, slice.Max(c.src))
			}
		})
	}
}

func TestMin(t *testing.T) {
	cases := []struct {
		title  string
		src    []int
		want   int
		panics bool
	}{
		{
			title:  "empty",
			src:    []int{},
			want:   0,
			panics: true,
		},
		{
			title:  "single",
			src:    []int{1},
			want:   1,
			panics: false,
		},
		{
			title:  "multiple",
			src:    []int{1, 2, 3, 4, 5},
			want:   1,
			panics: false,
		},
		{
			title:  "negative",
			src:    []int{-1, -2, -3, -4, -5},
			want:   -5,
			panics: false,
		},
		{
			title:  "mixed",
			src:    []int{-1, 2, -3, 4, -5},
			want:   -5,
			panics: false,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if c.panics {
				assert.Panics(t, func() { slice.Min(c.src) })
			} else {
				assert.Equal(t, c.want, slice.Min(c.src))
			}
		})
	}
}

func TestExtremeValue(t *testing.T) {
	cmp := func(a, b complex128) bool {
		return real(a)+imag(a) > real(b)+imag(b)
	}
	cases := []struct {
		title  string
		src    []complex128
		want   complex128
		panics bool
	}{
		{
			title:  "empty",
			src:    []complex128{},
			want:   0,
			panics: true,
		},
		{
			title:  "single",
			src:    []complex128{1},
			want:   1,
			panics: false,
		},
		{
			title:  "multiple",
			src:    []complex128{1 + 3i, 3 + 8i, 3 + 3i, 4 + 4i, 5 + 5i},
			want:   3 + 8i,
			panics: false,
		},
		{
			title:  "negative",
			src:    []complex128{-1 - 3i, -2 - 8i, -3 - 3i, -4 - 4i, -5 - 5i},
			want:   -1 - 3i,
			panics: false,
		},
		{
			title:  "mixed",
			src:    []complex128{-1 - 3i, 2 + 8i, -3 - 3i, 4 + 4i, -5 - 5i},
			want:   2 + 8i,
			panics: false,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if c.panics {
				assert.Panics(t, func() { slice.ExtremeValue(c.src, cmp) })
			} else {
				assert.Equal(t, c.want, slice.ExtremeValue(c.src, cmp))
			}
		})
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		title string
		src   []int
		want  int
	}{
		{
			title: "empty",
			src:   []int{},
			want:  0,
		},
		{
			title: "single",
			src:   []int{1},
			want:  1,
		},
		{
			title: "multiple",
			src:   []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			title: "negative",
			src:   []int{-1, -2, -3, -4, -5},
			want:  -15,
		},
		{
			title: "mixed",
			src:   []int{-1, 2, -3, 4, -5},
			want:  -3,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.want, slice.Sum(c.src))
		})
	}
}
