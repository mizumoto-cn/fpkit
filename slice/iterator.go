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
package slice

import "github.com/mizumoto-cn/fpkit/internal/slice"

type Iterator[T any] struct {
	src    []T
	cursor int
}

func NewIterator[T any](src []T) *Iterator[T] {
	return &Iterator[T]{src: src}
}

func (i *Iterator[T]) HasNext() bool {
	return i.cursor < len(i.src)
}

func (i *Iterator[T]) Next() T {
	v := i.src[i.cursor]
	i.cursor++
	return v
}

func (i *Iterator[T]) Remove() {
	i.src, _, _ = slice.Delete(i.src, i.cursor-1)
}

func (i *Iterator[T]) Slice() []T {
	return i.src
}

func (i *Iterator[T]) Head() T {
	return i.src[0]
}

func (i *Iterator[T]) Tail() []T {
	return i.src[1:]
}

func (i *Iterator[T]) Last() T {
	return i.src[len(i.src)-1]
}

func (i *Iterator[T]) Init() []T {
	return i.src[:len(i.src)-1]
}

func (i *Iterator[T]) Index() int {
	return i.cursor - 1
}

func (i *Iterator[T]) Reset() {
	i.cursor = 0
}
