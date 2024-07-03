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
package functional

import (
	"sync"
)

// Curry2 returns a curried version of a function that takes 2 arguments.
//
// Input: fn func(A, B) R
// Output: func(A) func(B) R
//
// Example:
//
// func Add(a, b int) int { return a + b }
//
// Add2 := Curry2(Add)
//
// Add2(1)(2) // 3
func Curry2[A, B, R any](fn func(A, B) R) func(A) func(B) R {
	return func(a A) func(B) R {
		return func(b B) R {
			return fn(a, b)
		}
	}
}

// Curry3 returns a curried version of a function that takes 3 arguments.
func Curry3[A, B, C, R any](fn func(A, B, C) R) func(A) func(B) func(C) R {
	return func(a A) func(B) func(C) R {
		return func(b B) func(C) R {
			return func(c C) R {
				return fn(a, b, c)
			}
		}
	}
}

// Curry4 returns a curried version of a function that takes 4 arguments.
func Curry4[A, B, C, D, R any](fn func(A, B, C, D) R) func(A) func(B) func(C) func(D) R {
	return func(a A) func(B) func(C) func(D) R {
		return func(b B) func(C) func(D) R {
			return func(c C) func(D) R {
				return func(d D) R {
					return fn(a, b, c, d)
				}
			}
		}
	}
}

// Curry5 returns a curried version of a function that takes 5 arguments.
func Curry5[A, B, C, D, E, R any](fn func(A, B, C, D, E) R) func(A) func(B) func(C) func(D) func(E) R {
	return func(a A) func(B) func(C) func(D) func(E) R {
		return func(b B) func(C) func(D) func(E) R {
			return func(c C) func(D) func(E) R {
				return func(d D) func(E) R {
					return func(e E) R {
						return fn(a, b, c, d, e)
					}
				}
			}
		}
	}
}

// Curry6 returns a curried version of a function that takes 6 arguments.
func Curry6[A, B, C, D, E, F, R any](fn func(A, B, C, D, E, F) R) func(A) func(B) func(C) func(D) func(E) func(F) R {
	return func(a A) func(B) func(C) func(D) func(E) func(F) R {
		return func(b B) func(C) func(D) func(E) func(F) R {
			return func(c C) func(D) func(E) func(F) R {
				return func(d D) func(E) func(F) R {
					return func(e E) func(F) R {
						return func(f F) R {
							return fn(a, b, c, d, e, f)
						}
					}
				}
			}
		}
	}
}

// CurryDef defines a curried function type.
type CurryDef[T any, R any] struct {
	fn     func(c *CurryDef[T, R], args ...T) R
	result R
	isDone AtomicBool

	callM sync.Mutex
	args  []T
}

// CurryNew creates a new Curry instance.
func CurryNew(fn func(c *CurryDef[any, any], args ...any) any) *CurryDef[any, any] {
	return CurryNewGenerics(fn)
}

// CurryNewGenerics creates a new Curry instance with generics.
func CurryNewGenerics[T any, R any](fn func(c *CurryDef[T, R], args ...T) R) *CurryDef[T, R] {
	c := &CurryDef[T, R]{fn: fn}
	return c
}

// Call calls the curried function with arguments (partly or fully).
func (currySelf *CurryDef[T, R]) Call(args ...T) *CurryDef[T, R] {
	currySelf.callM.Lock()
	if !currySelf.isDone.Get() {
		currySelf.args = append(currySelf.args, args...)
		currySelf.result = currySelf.fn(currySelf, currySelf.args...)
	}
	currySelf.callM.Unlock()
	return currySelf
}

// MarkDone marks the curried function as done.
func (currySelf *CurryDef[T, R]) MarkDone() {
	currySelf.isDone.Set(true)
}

// IsDone checks if the curried function is done.
func (currySelf *CurryDef[T, R]) IsDone() bool {
	return currySelf.isDone.Get()
}

// Result returns the result of the curried function.
func (currySelf *CurryDef[T, R]) Result() R {
	return currySelf.result
}

// Curry creates a curried function.
// var Curry *CurryDef[any, any]
