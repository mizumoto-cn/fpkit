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

// Tests for SliceOf
func TestSliceOf(t *testing.T) {
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
			got := functional.SliceOf(c.v...)
			assert.True(t, reflect.DeepEqual(got, c.want))
		})
	}
}

// Tests for IsPtr
func TestIsPtr(t *testing.T) {
	vint := 1
	vptr := &vint
	cases := []struct {
		title string
		v     any
		want  bool
	}{
		{
			title: "int",
			v:     vint,
			want:  false,
		},
		{
			title: "pointer",
			v:     vptr,
			want:  true,
		},
		{
			title: "nil",
			v:     nil,
			want:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got := functional.IsPtr(c.v)
			assert.Equal(t, c.want, got)
		})
	}
}

// Tests for Kind
func TestKind(t *testing.T) {
	cases := []struct {
		title string
		v     any
		want  reflect.Kind
	}{
		{
			title: "int",
			v:     1,
			want:  reflect.Int,
		},
		{
			title: "string",
			v:     "hello",
			want:  reflect.String,
		},
		{
			title: "struct",
			v:     struct{ Name string }{"mizumoto"},
			want:  reflect.Struct,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got := functional.Kind(c.v)
			assert.Equal(t, c.want, got)
		})
	}
}

// Tests for IsNil
func TestIsNil(t *testing.T) {
	var vptr *int
	vint := 1
	cases := []struct {
		title string
		v     any
		want  bool
	}{
		{
			title: "nil pointer",
			v:     vptr,
			want:  true,
		},
		{
			title: "non-nil pointer",
			v:     &vint,
			want:  false,
		},
		{
			title: "non-pointer",
			v:     vint,
			want:  false,
		},
		{
			title: "invalid value",
			v:     nil,
			want:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got := functional.IsNil(c.v)
			assert.Equal(t, c.want, got)
		})
	}
}
