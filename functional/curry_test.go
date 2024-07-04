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

func TestCurry2(t *testing.T) {
	add := func(a, b int) int { return a + b }
	add2 := functional.Curry2(add)
	assert.Equal(t, 3, add2(1)(2))
}

func TestCurry3(t *testing.T) {
	add := func(a, b, c int) int { return a + b + c }
	add3 := functional.Curry3(add)
	assert.Equal(t, 6, add3(1)(2)(3))
}

func TestCurry4(t *testing.T) {
	add := func(a, b, c, d int) int { return a + b + c + d }
	add4 := functional.Curry4(add)
	assert.Equal(t, 10, add4(1)(2)(3)(4))
}

func TestCurry5(t *testing.T) {
	add := func(a, b, c, d, e int) int { return a + b + c + d + e }
	add5 := functional.Curry5(add)
	assert.Equal(t, 15, add5(1)(2)(3)(4)(5))
}

func TestCurry6(t *testing.T) {
	add := func(a, b, c, d, e, f int) int { return a + b + c + d + e + f }
	add6 := functional.Curry6(add)
	assert.Equal(t, 21, add6(1)(2)(3)(4)(5)(6))
}

func TestCurry(t *testing.T) {
	addFunc := func(cDef *functional.CurryDef[any, any], args ...any) any {
		if len(args) < 3 {
			return cDef
		}
		a, b, c := args[0].(int), args[1].(int), args[2].(int)
		cDef.MarkDone() // Mark the curry chain as done
		return a + b + c
	}

	// Create a new curry function
	Curry := functional.CurryNew(addFunc)

	// Call the curry function with 3 arguments
	result := Curry.Call(1).Call(2).Call(3).Result().(int)
	assert.Equal(t, 6, result)
}
