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
	"sync/atomic"
	"unsafe"

	"github.com/mizumoto-cn/fpkit/internal/err"
)

// LinkedQueue is a queue implemented using a linked list.
// LinkedQueue is thread-safe.
type LinkedQueue[T any] struct {
	// head *node[T]
	head unsafe.Pointer
	// tail *node[T]
	tail unsafe.Pointer
	size int32
}

// node is a node in the linked list.
type node[T any] struct {
	value T
	next  unsafe.Pointer
}

var _ Queue[int] = (*LinkedQueue[int])(nil)

// NewLinkedQueue creates a new LinkedQueue.
func NewLinkedQueue[T any]() *LinkedQueue[T] {
	n := unsafe.Pointer(&node[T]{})
	return &LinkedQueue[T]{head: n, tail: n}
}

// cas: compare-and-swap atomic operation.
func cas(ptr *unsafe.Pointer, old, new unsafe.Pointer) bool {
	return atomic.CompareAndSwapPointer(ptr, old, new)
}

// updateSize updates the size of the queue atomically.
func (q *LinkedQueue[T]) updateSize(delta int32) {
	atomic.AddInt32(&q.size, delta)
}

// Push adds an element to the end of the queue.
func (q *LinkedQueue[T]) Push(t T) error {
	n := &node[T]{value: t}
	ptr := unsafe.Pointer(n)
	for {
		tailPtr := atomic.LoadPointer(&q.tail)
		tail := (*node[T])(tailPtr)
		next := atomic.LoadPointer(&tail.next)
		if next != nil {
			// Someone has edited the tail
			// Simply retry to wait for the q.tail to be updated
			continue
		}
		if cas(&tail.next, next, ptr) {
			cas(&q.tail, tailPtr, ptr)
			// atomic.AddInt32(&q.size, 1)
			q.updateSize(1)
			return nil
		}
	}
}

// Pop removes and returns the first element in the queue.
func (q *LinkedQueue[T]) Pop() (T, error) {
	for {
		hp := atomic.LoadPointer(&q.head)
		head := (*node[T])(hp)
		tp := atomic.LoadPointer(&q.tail)
		tail := (*node[T])(tp)
		if head == tail {
			// The queue is empty
			// Or at least we think it is empty
			var zero T
			return zero, err.ErrEmptyQueue
		}
		np := atomic.LoadPointer(&head.next)
		if cas(&q.head, hp, np) {
			next := (*node[T])(np)
			q.updateSize(-1)
			return next.value, nil
		}
	}
}

// Empty returns true if the queue is empty.
func (q *LinkedQueue[T]) Empty() bool {
	size := atomic.LoadInt32(&q.size)
	return size == 0
}

// Size returns the number of elements in the queue.
func (q *LinkedQueue[T]) Size() int {
	return (int)(atomic.LoadInt32(&q.size))
}

// Cap returns the capacity of the queue.
// For boundless queues, Cap returns -1.
func (q *LinkedQueue[T]) Cap() int {
	return -1
}

// Back returns the last element in the queue.
// Deprecated: Not necessarily real-time as the queue may be updated
func (q *LinkedQueue[T]) Back() (T, error) {
	if q.Empty() {
		var zero T
		return zero, err.ErrEmptyQueue
	}
	tp := atomic.LoadPointer(&q.tail)
	tail := (*node[T])(tp)
	return tail.value, nil
}

// Front returns the first element in the queue.
func (q *LinkedQueue[T]) Front() (T, error) {
	if q.Empty() {
		var zero T
		return zero, err.ErrEmptyQueue
	}
	hp := atomic.LoadPointer(&q.head)
	head := (*node[T])(hp)
	np := atomic.LoadPointer(&head.next)
	next := (*node[T])(np)
	return next.value, nil
}

// Clear removes all elements from the queue.
// We do deprecate this method to be used in situations where
// the queue is foreseeable to be used soon again after being cleared,
// especially in a multi-threaded environment.
//
//	TODO: may need updates on this method as it surely is error-prone
//	and may cause unexpected behaviors.
func (q *LinkedQueue[T]) Clear() error {
	for !q.Empty() {
		_, err := q.Pop()
		if err != nil {
			return err
		}
	}
	return nil
}

// Slice returns a slice of the elements in the queue.
// Deprecated: Not necessarily real-time as the queue may be updated
func (q *LinkedQueue[T]) Slice() []T {
	if q.Empty() {
		return nil
	}
	hp := atomic.LoadPointer(&q.head)
	head := (*node[T])(hp)
	np := atomic.LoadPointer(&head.next)
	next := (*node[T])(np)
	slice := make([]T, 0, q.Size())
	for next != nil {
		slice = append(slice, next.value)
		np = atomic.LoadPointer(&next.next)
		next = (*node[T])(np)
	}
	return slice
}
