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

// Compose returns a function that is the composition of the given functions from right to left.
func Compose[T any](fns ...func(T) T) func(T) T {
	return func(x T) T {
		for i := len(fns) - 1; i >= 0; i-- {
			x = fns[i](x)
		}
		return x
	}
}

// Pipe returns a function that is the composition of the given functions from left to right.
func Pipe[T any](fns ...func(T) T) func(T) T {
	return func(x T) T {
		for i := range fns {
			x = fns[i](x)
		}
		return x
	}
}
