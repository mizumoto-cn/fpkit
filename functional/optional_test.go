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
)

func TestJust(t *testing.T) {
	opt := functional.Just(42)
	if !opt.IsPresent() {
		t.Error("Expected IsPresent to be true")
	}
	if opt.IsNil() {
		t.Error("Expected IsNil to be false")
	}
	if !opt.IsValid() {
		t.Error("Expected IsValid to be true")
	}
	if opt.IsPtr() {
		t.Error("Expected IsPtr to be false")
	}
	if opt.Unwrap() != 42 {
		t.Errorf("Expected Unwrap to return 42, got %v", opt.Unwrap())
	}
}

func TestNone(t *testing.T) {
	if functional.None.IsPresent() {
		t.Error("Expected IsPresent to be false")
	}
	if !functional.None.IsNil() {
		t.Error("Expected IsNil to be true")
	}
	if functional.None.IsValid() {
		t.Error("Expected IsValid to be false")
	}
	if functional.None.IsPtr() {
		t.Error("Expected IsPtr to be false")
	}
	if functional.None.UnwrapAny() != nil {
		t.Errorf("Expected UnwrapAny to return nil, got %v", functional.None.UnwrapAny())
	}
}

func TestOrElse(t *testing.T) {
	opt := functional.Just(42)
	if opt.OrElse(0) != 42 {
		t.Errorf("Expected OrElse to return 42, got %v", opt.OrElse(0))
	}
	if functional.Maybe.Just(nil).OrElse(0) != 0 {
		t.Errorf("Expected OrElse to return 0, got %v", functional.Maybe.Just(nil).OrElse(0))
	}
	if functional.None.OrElse(0) != 0 {
		t.Errorf("Expected OrElse to return 0, got %v", functional.None.OrElse(0))
	}
}

func TestClone(t *testing.T) {
	opt := functional.Just(42)
	clone := opt.Clone()
	if !clone.IsPresent() || clone.Unwrap() != 42 {
		t.Error("Expected clone to be present and equal to 42")
	}

	noneClone := functional.None.Clone()
	if noneClone.IsPresent() {
		t.Error("Expected noneClone to be not present")
	}
}

func TestFlatMap(t *testing.T) {
	opt := functional.Just(10)
	result := opt.FlatMap(func(v int) functional.Optional[int] {
		return functional.Just(v / 2)
	})
	if !result.IsPresent() || result.Unwrap() != 5 {
		t.Errorf("Expected FlatMap result to be present and equal to 5, got %v", result.Unwrap())
	}

	noneResult := functional.None.FlatMap(func(v any) functional.Optional[any] {
		return functional.None.Just(nil)
	})
	if noneResult.IsPresent() {
		t.Error("Expected noneResult to be not present")
	}
}

func TestIfPresent(t *testing.T) {
	opt := functional.Just(10)
	var called bool
	opt.IfPresent(func() {
		called = true
	})
	if !called {
		t.Error("Expected IfPresent to call the function")
	}

	called = false
	functional.None.IfPresent(func() {
		called = true
	})
	if called {
		t.Error("Expected IfPresent to not call the function")
	}
}

func TestKindAndType(t *testing.T) {
	opt := functional.Just(10)
	if opt.Kind() != reflect.Int {
		t.Errorf("Expected Kind to be reflect.Int, got %v", opt.Kind())
	}
	if opt.Type() != reflect.TypeOf(10) {
		t.Errorf("Expected Type to be reflect.TypeOf(int), got %v", opt.Type())
	}

	if functional.None.Kind() != reflect.Invalid {
		t.Errorf("Expected Kind to be reflect.Invalid, got %v", functional.None.Kind())
	}
	if functional.None.Type() != reflect.TypeOf(nil) {
		t.Errorf("Expected Type to be reflect.TypeOf(nil), got %v", functional.None.Type())
	}
}

func TestIsKindOfAndIsTypeOf(t *testing.T) {
	opt := functional.Just(10)
	if !opt.IsKindOf(reflect.Int) {
		t.Errorf("Expected IsKindOf(reflect.Int) to be true")
	}
	if !opt.IsTypeOf(reflect.TypeOf(10)) {
		t.Errorf("Expected IsTypeOf(reflect.TypeOf(10)) to be true")
	}

	if functional.None.IsKindOf(reflect.Int) {
		t.Errorf("Expected IsKindOf(reflect.Int) to be false")
	}
	if functional.None.IsTypeOf(reflect.TypeOf(10)) {
		t.Errorf("Expected IsTypeOf(reflect.TypeOf(10)) to be false")
	}
}

func TestMaybeJust(t *testing.T) {
	opt := functional.Maybe.Just(42)
	if !opt.IsPresent() {
		t.Error("Expected IsPresent to be true")
	}
	if opt.IsNil() {
		t.Error("Expected IsNil to be false")
	}
	if opt.UnwrapAny() != 42 {
		t.Errorf("Expected UnwrapAny to return 42, got %v", opt.UnwrapAny())
	}
}

func TestMakeCloneNil(t *testing.T) {
	opt := functional.Maybe.Just(nil)
	clone := opt.Clone()
	if clone.IsPresent() {
		t.Error("Expected clone not to be present")
	}
	if clone.UnwrapAny() != nil {
		t.Errorf("Expected clone to be nil, got %v", clone.UnwrapAny())
	}
}

func TestMakeClonePtr(t *testing.T) {
	opt := functional.Maybe.Just(&struct{}{})
	clone := opt.Clone()
	if !clone.IsPresent() {
		t.Error("Expected clone to be present")
	}
	if clone.UnwrapAny() == nil {
		t.Error("Expected clone not to be nil")
	}
}
