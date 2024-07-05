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

func add1(x int) int {
	return x + 1
}

func multiplyBy5(x int) int {
	return x * 5
}

func TestCompose(t *testing.T) {
	cases := []struct {
		name     string
		fns      []func(int) int
		input    int
		expected int
	}{
		{
			name:     "multiply by 5 then add 1", // right to left
			fns:      []func(int) int{add1, multiplyBy5},
			input:    1,
			expected: 6,
		},
		{
			name:     "add 1 then multiply by 5",
			fns:      []func(int) int{multiplyBy5, add1},
			input:    1,
			expected: 10,
		},
		{
			name:     "add 1 only",
			fns:      []func(int) int{add1},
			input:    1,
			expected: 2,
		},
		{
			name:     "do nothing",
			fns:      []func(int) int{},
			input:    1,
			expected: 1,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := functional.Compose(tc.fns...)(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPipe(t *testing.T) {
	cases := []struct {
		name     string
		fns      []func(int) int
		input    int
		expected int
	}{
		{
			name:     "add 1 then multiply by 5", // left to right
			fns:      []func(int) int{add1, multiplyBy5},
			input:    1,
			expected: 10,
		},
		{
			name:     "multiply by 5 then add 1",
			fns:      []func(int) int{multiplyBy5, add1},
			input:    1,
			expected: 6,
		},
		{
			name:     "add 1 only",
			fns:      []func(int) int{add1},
			input:    1,
			expected: 2,
		},
		{
			name:     "do nothing",
			fns:      []func(int) int{},
			input:    1,
			expected: 1,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := functional.Pipe(tc.fns...)(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
