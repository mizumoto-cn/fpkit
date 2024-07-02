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

func TestContains(t *testing.T) {
	cases := []struct {
		title string
		src   []int
		value int
		want  bool
	}{
		{
			title: "empty",
			src:   []int{},
			value: 1,
			want:  false,
		},
		{
			title: "single",
			src:   []int{1},
			value: 1,
			want:  true,
		},
		{
			title: "multiple",
			src:   []int{1, 2, 3, 4, 5},
			value: 3,
			want:  true,
		},
		{
			title: "negative",
			src:   []int{-1, -2, -3, -4, -5},
			value: -3,
			want:  true,
		},
		{
			title: "mixed",
			src:   []int{-1, 2, -3, 4, -5},
			value: 4,
			want:  true,
		},
		{
			title: "not found",
			src:   []int{1, 2, 3, 4, 5},
			value: 6,
			want:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.want, slice.Contains(c.src, c.value))
		})
	}
}

func TestContainsAny(t *testing.T) {
	cases := []struct {
		title  string
		src    []int
		values []int
		want   bool
	}{
		{
			title:  "empty",
			src:    []int{},
			values: []int{1, 2, 3},
			want:   false,
		},
		{
			title:  "single",
			src:    []int{1},
			values: []int{1, 2, 3},
			want:   true,
		},
		{
			title:  "multiple",
			src:    []int{1, 2, 3, 4, 5},
			values: []int{3, 4, 5},
			want:   true,
		},
		{
			title:  "negative",
			src:    []int{-1, -2, -3, -4, -5},
			values: []int{-3, -4, -5},
			want:   true,
		},
		{
			title:  "mixed",
			src:    []int{-1, 2, -3, 4, -5},
			values: []int{4, 5, 6},
			want:   true,
		},
		{
			title:  "not found",
			src:    []int{1, 2, 3, 4, 5},
			values: []int{6, 7, 8},
			want:   false,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.want, slice.ContainsAny(c.src, c.values...))
		})
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		title  string
		src    []int
		values []int
		want   bool
	}{
		{
			title:  "empty",
			src:    []int{},
			values: []int{1, 2, 3},
			want:   false,
		},
		{
			title:  "single",
			src:    []int{1},
			values: []int{1},
			want:   true,
		},
		{
			title:  "multiple",
			src:    []int{1, 2, 3, 4, 5},
			values: []int{3, 4, 5},
			want:   true,
		},
		{
			title:  "negative",
			src:    []int{-1, -2, -3, -4, -5},
			values: []int{-3, -4, -5},
			want:   true,
		},
		{
			title:  "mixed",
			src:    []int{-1, 2, -3, 4, -5},
			values: []int{4, -5},
			want:   true,
		},
		{
			title:  "not found",
			src:    []int{1, 2, 3, 4, 5},
			values: []int{6, 7, 8},
			want:   false,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.want, slice.ContainsAll(c.src, c.values))
		})
	}
}
