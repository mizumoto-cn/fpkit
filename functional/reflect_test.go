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
	"reflect"
	"testing"

	"github.com/mizumoto-cn/fpkit/functional"

	"github.com/stretchr/testify/assert"
)

// . Tests for PtrOf
func TestPtrOf(t *testing.T) {
	vint := 1
	vstring := "hello"
	vstruct := struct{ Name string }{Name: "mizumoto"}
	cases := []struct {
		title string
		v     any
		want  any
	}{
		{
			title: "int",
			v:     vint,
			want:  vint,
		},
		{
			title: "string",
			v:     vstring,
			want:  vstring,
		},
		{
			title: "struct",
			v:     vstruct,
			want:  vstruct,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got := functional.PtrOf(c.v)
			// To properly test the PtrOf function, you should compare the values pointed to by the pointers for equality, not the pointers themselves.
			assert.True(t, reflect.DeepEqual(reflect.ValueOf(got).Elem().Interface(), c.want))
		})
	}
}
