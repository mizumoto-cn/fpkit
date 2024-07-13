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
	"github.com/mizumoto-cn/fpkit/functional"
	"github.com/mizumoto-cn/fpkit/internal/err"
)

type PriorityQueue[T any] struct {
	cmp   functional.ComparatorAny[T]
	data  []T
	cap   int
	head  int
	tail  int
	count int
}

var _ Queue[int] = (*PriorityQueue[int])(nil)

func NewPriorityQueue[T any](cmp functional.ComparatorAny[T], cap int) (*PriorityQueue[T], error) {
	if cap <= 0 {
		return nil, err.NewQueueCapacityError(cap)
	}
	return &PriorityQueue[T]{
		cmp:   cmp,
		data:  make([]T, cap),
		cap:   cap,
		head:  0,
		tail:  0,
		count: 0,
	}, nil
}

func (pq *PriorityQueue[T]) Push(v T) error {
	if pq.count == pq.cap {
		return err.NewQueueFullError(pq.cap, pq.count)
	}
	pq.data[pq.tail] = v
	pq.tail = (pq.tail + 1) % pq.cap
	pq.count++
	pq.up((pq.tail - 1 + pq.cap) % pq.cap)
	return nil
}

func (pq *PriorityQueue[T]) up(i int) {
	for {
		parent := (i - 1 + pq.cap) % pq.cap
		if i == pq.head || pq.cmp(pq.data[parent], pq.data[i]) {
			break
		}
		pq.data[parent], pq.data[i] = pq.data[i], pq.data[parent]
		i = parent
	}
}

func (pq *PriorityQueue[T]) down(i int) {
	for {
		left := (2*i + 1) % pq.cap
		if left == pq.tail || left == pq.head {
			break
		}
		j := left
		right := (left + 1) % pq.cap
		if right != pq.tail && right != pq.head && pq.cmp(pq.data[right], pq.data[left]) {
			j = right
		}
		if pq.cmp(pq.data[i], pq.data[j]) {
			break
		}
		pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
		i = j
	}
}

func (pq *PriorityQueue[T]) Pop() (T, error) {
	if pq.count == 0 {
		var zero T
		return zero, err.NewIndexOutOfRangeError(0, pq.count)
	}
	v := pq.data[pq.head]
	pq.head = (pq.head + 1) % pq.cap
	pq.count--
	pq.down(pq.head)
	return v, nil
}

func (pq *PriorityQueue[T]) Front() (T, error) {
	if pq.count == 0 {
		var zero T
		return zero, err.NewIndexOutOfRangeError(0, pq.count)
	}
	return pq.data[pq.head], nil
}

func (pq *PriorityQueue[T]) Back() (T, error) {
	if pq.count == 0 {
		var zero T
		return zero, err.NewIndexOutOfRangeError(0, pq.count)
	}
	tail := (pq.tail - 1 + pq.cap) % pq.cap
	return pq.data[tail], nil
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.count
}

func (pq *PriorityQueue[T]) Cap() int {
	return pq.cap
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.count == 0
}

func (pq *PriorityQueue[T]) Full() bool {
	return pq.count == pq.cap
}

func (pq *PriorityQueue[T]) Clear() error {
	pq.head = 0
	pq.tail = 0
	pq.count = 0
	return nil
}
