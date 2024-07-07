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
package queue

import "github.com/mizumoto-cn/fpkit/internal/err"

// NewBasicQueue creates a new basicQueue.
func NewBasicQueue[T any](cap int) Queue[T] {
	return &basicQueue[T]{data: make([]T, 0, cap), cap: cap}
}

// basicQueue is a basic generic FIFO queue.
type basicQueue[T any] struct {
	data []T
	cap  int
}

// Push adds an element to the end of the queue.
func (q *basicQueue[T]) Push(t T) error {
	// check if the queue is full
	if len(q.data) >= cap(q.data) {
		return err.NewQueueFullError()
	}
}
