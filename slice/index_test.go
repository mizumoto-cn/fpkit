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

func TestIndex(t *testing.T) {
	cases := []struct {
		title    string
		src      []int
		value    int
		expected int
	}{
		{
			title:    "value present at beginning",
			src:      []int{1, 2, 3, 4},
			value:    1,
			expected: 0,
		},
		{
			title:    "value present in middle",
			src:      []int{1, 2, 3, 4},
			value:    3,
			expected: 2,
		},
		{
			title:    "value not present",
			src:      []int{1, 2, 3, 4},
			value:    5,
			expected: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.Index(c.src, c.value)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestIndexAll(t *testing.T) {
	var empty []int
	cases := []struct {
		title    string
		src      []int
		value    int
		expected []int
	}{
		{
			title:    "multiple occurrences",
			src:      []int{1, 2, 3, 1, 1},
			value:    1,
			expected: []int{0, 3, 4},
		},
		{
			title:    "no occurrences",
			src:      []int{1, 2, 3, 4},
			value:    5,
			expected: empty,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.IndexAll(c.src, c.value)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestLastIndex(t *testing.T) {
	cases := []struct {
		title    string
		src      []int
		value    int
		expected int
	}{
		{
			title:    "value present at end",
			src:      []int{1, 2, 3, 4},
			value:    4,
			expected: 3,
		},
		{
			title:    "value present in middle",
			src:      []int{1, 4, 3, 4},
			value:    4,
			expected: 3,
		},
		{
			title:    "value not present",
			src:      []int{1, 2, 3, 4},
			value:    5,
			expected: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.LastIndex(c.src, c.value)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestIndexMatchFunc(t *testing.T) {
	cases := []struct {
		title    string
		src      []int
		match    func(int) bool
		expected int
	}{
		{
			title:    "match found",
			src:      []int{1, 2, 3, 4},
			match:    func(i int) bool { return i == 3 },
			expected: 2,
		},
		{
			title:    "match not found",
			src:      []int{1, 2, 3, 4},
			match:    func(i int) bool { return i == 5 },
			expected: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.IndexMatchFunc(c.src, c.match)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestIndexAllMatchFunc(t *testing.T) {
	var empty []int
	cases := []struct {
		title    string
		src      []int
		match    func(int) bool
		expected []int
	}{
		{
			title:    "multiple matches",
			src:      []int{1, 2, 3, 1, 1},
			match:    func(i int) bool { return i == 1 },
			expected: []int{0, 3, 4},
		},
		{
			title:    "no matches",
			src:      []int{1, 2, 3, 4},
			match:    func(i int) bool { return i == 5 },
			expected: empty,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.IndexAllMatchFunc(c.src, c.match)
			assert.Equal(t, c.expected, result)
		})
	}
}
