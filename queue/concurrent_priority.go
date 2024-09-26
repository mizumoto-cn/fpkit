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
	"sync"

	"github.com/mizumoto-cn/fpkit/functional"
)

// ConcurrentPriorityQueue is a priority queue with a fixed capacity.
// It is thread-safe.
//
//	cpq := queue.NewConcurrentPriorityQueue[int](functional.Less[int], 3)
type ConcurrentPriorityQueue[T any] struct {
	pq PriorityQueue[T]
	m  sync.RWMutex
}

var _ Queue[int] = (*ConcurrentPriorityQueue[int])(nil)

// NewConcurrentPriorityQueue returns a new concurrent priority queue with the given capacity.
// The capacity must be greater than 0, otherwise it will return an error.
//
//	NewConcurrentPriorityQueue(func(a, b int) bool { return a < b }, 3)
func NewConcurrentPriorityQueue[T any](cmp functional.ComparatorAny[T], cap int) (*ConcurrentPriorityQueue[T], error) {
	pq, err := NewPriorityQueue[T](cmp, cap)
	if err != nil {
		return nil, err
	}
	return &ConcurrentPriorityQueue[T]{pq: *pq}, nil
}

// Push adds an element to the concurrent priority queue.
// If the queue is full, it will block until there is space.
func (cpq *ConcurrentPriorityQueue[T]) Push(v T) error {
	cpq.m.Lock()
	defer cpq.m.Unlock()
	return cpq.pq.Push(v)
}

// Pop removes and returns the element with the highest priority from the concurrent priority queue.
func (cpq *ConcurrentPriorityQueue[T]) Pop() (T, error) {
	cpq.m.Lock()
	defer cpq.m.Unlock()
	return cpq.pq.Pop()
}

// Front returns the element with the highest priority from the concurrent priority queue.
func (cpq *ConcurrentPriorityQueue[T]) Front() (T, error) {
	cpq.m.RLock()
	defer cpq.m.RUnlock()
	return cpq.pq.Front()
}

// Back returns the element with the lowest priority from the concurrent priority queue.
func (cpq *ConcurrentPriorityQueue[T]) Back() (T, error) {
	cpq.m.RLock()
	defer cpq.m.RUnlock()
	return cpq.pq.Back()
}

// Size returns the number of elements in the concurrent priority queue.
func (cpq *ConcurrentPriorityQueue[T]) Size() int {
	cpq.m.RLock()
	defer cpq.m.RUnlock()
	return cpq.pq.Size()
}

// Cap returns the capacity of the concurrent priority queue.
func (cpq *ConcurrentPriorityQueue[T]) Cap() int {
	cpq.m.RLock()
	defer cpq.m.RUnlock()
	return cpq.pq.Cap()
}

// Empty returns true if the concurrent priority queue is empty.
func (cpq *ConcurrentPriorityQueue[T]) Empty() bool {
	cpq.m.RLock()
	defer cpq.m.RUnlock()
	return cpq.pq.Empty()
}

// Full returns true if the concurrent priority queue is full.
func (cpq *ConcurrentPriorityQueue[T]) Full() bool {
	cpq.m.RLock()
	defer cpq.m.RUnlock()
	return cpq.pq.Full()
}

// Clear removes all elements from the concurrent priority queue.
func (cpq *ConcurrentPriorityQueue[T]) Clear() error {
	cpq.m.Lock()
	defer cpq.m.Unlock()
	return cpq.pq.Clear()
}
