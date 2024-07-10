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

import (
	"github.com/mizumoto-cn/fpkit/internal/err"
)

type BasicQueue[T any] struct {
	data []T
	head int
	tail int // tail points to the next available slot
	size int
	cap  int
}

func NewBasicQueue[T any](cap int) *BasicQueue[T] {
	return &BasicQueue[T]{data: make([]T, cap), cap: cap}
}

// Push adds an element to the end of the queue.
func (q *BasicQueue[T]) Push(t T) error {
	// check if the queue is full
	if q.size >= q.cap {
		return err.NewQueueFullError(q.cap, q.size)
	}
	q.data[q.tail%q.cap] = t
	q.tail = (q.tail + 1) % q.cap
	q.size++
	return nil
}

// Pop removes and returns the first element in the queue.
func (q *BasicQueue[T]) Pop() (T, error) {
	var zero T
	// check if the queue is empty
	if q.size == 0 {
		return zero, err.NewIndexOutOfRangeError(0, 0)
	}
	t := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	q.size--
	return t, nil
}

// Empty returns true if the queue is empty.
func (q *BasicQueue[T]) Empty() bool {
	return q.size == 0
}

// Size returns the number of elements in the queue.
func (q *BasicQueue[T]) Size() int {
	return q.size
}

// Front returns the first element in the queue.
func (q *BasicQueue[T]) Front() (T, error) {
	var zero T
	// check if the queue is empty
	if q.size == 0 {
		return zero, err.NewIndexOutOfRangeError(0, 0)
	}
	return q.data[q.head], nil
}

// Back returns the last element in the queue.
func (q *BasicQueue[T]) Back() (T, error) {
	var zero T
	// check if the queue is empty
	if q.size == 0 {
		return zero, err.NewIndexOutOfRangeError(0, 0)
	}
	// In case of tail == 0, we need to wrap around to the end of the slice.
	return q.data[(q.tail-1+q.cap)%q.cap], nil
}

// Resize changes the capacity of the queue.
// if shrink is true, the queue will shrink to the new capacity if the new capacity is less than the current size.
// else the queue will keep the capacity greater than or equal to the current size.
func (q *BasicQueue[T]) Resize(cap int, shrink bool) error {
	if cap < 0 {
		return err.NewQueueResizeError(cap, 0)
	}
	if shrink {
		if cap < q.size {
			q.size = cap
		}
	} else {
		if cap < q.size {
			return err.NewQueueResizeError(cap, q.size)
		}
	}

	t := make([]T, cap)
	for i := 0; i < q.size; i++ {
		t[i] = q.data[(q.head+i)%q.cap]
	}
	q.data = t
	q.head = 0
	q.tail = q.size
	q.cap = cap
	return nil
}

// Cap returns the capacity of the queue.
func (q *BasicQueue[T]) Cap() int {
	return q.cap
}

// Clear removes all elements in the queue.
func (q *BasicQueue[T]) Clear() error {
	q.head = 0
	q.tail = 0
	q.size = 0
	return nil
}

// Swap swaps the contents of two queues.
func (q *BasicQueue[T]) Swap(q2 Queue[T]) {
	qq2 := q2.(*BasicQueue[T])
	q.data, qq2.data = qq2.data, q.data
	q.head, qq2.head = qq2.head, q.head
	q.tail, qq2.tail = qq2.tail, q.tail
	q.size, qq2.size = qq2.size, q.size
	q.cap, qq2.cap = qq2.cap, q.cap
}
