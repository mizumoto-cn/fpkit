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
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

// ArrayBlockingQueue is a thread-safe bounded queue.
type ArrayBlockingQueue[T any] struct {
	items    []T
	cap      int
	head     int
	tail     int
	size     int
	lock     sync.Mutex
	notEmpty *semaphore.Weighted
	notFull  *semaphore.Weighted
}

// NewArrayBlockingQueue creates a new ArrayBlockingQueue.
func NewArrayBlockingQueue[T any](cap int) *ArrayBlockingQueue[T] {
	notEmpty := semaphore.NewWeighted(int64(cap))
	// Acquire all slots in the notEmpty semaphore.
	// Not sure if it is necessary to panic if _ is actually an error.
	_ = notEmpty.Acquire(context.Background(), int64(cap))
	return &ArrayBlockingQueue[T]{
		items:    make([]T, cap),
		cap:      cap,
		notEmpty: notEmpty,
		notFull:  semaphore.NewWeighted(int64(cap)),
	}
}

// Push adds an element to the queue.
// When cancelled or timeout, return context.Canceled or context.DeadlineExceeded.
// Shall always use errors.Is(err, context.Canceled) or errors.Is(err, context.DeadlineExceeded) to check the error.
func (q *ArrayBlockingQueue[T]) Push(ctx context.Context, t T) error {
	// Acquire a slot in the semaphore, blocking if necessary.
	if err := q.notFull.Acquire(ctx, 1); err != nil {
		return err
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	// Check if the context has already been cancelled when the lock is acquired.
	if ctx.Err() != nil {
		// Release a slot for the notFull semaphore.
		q.notFull.Release(1)
		return ctx.Err()
	}

	q.items[q.tail] = t
	q.tail = (q.tail + 1) % q.cap
	q.size++

	// Release a slot for the notEmpty semaphore.
	q.notEmpty.Release(1)

	return nil
}

// TryPop removes and returns an element from the queue.
// When cancelled or timeout, return context.Canceled or context.DeadlineExceeded.
// Shall always use errors.Is(err, context.Canceled) or errors.Is(err, context.DeadlineExceeded) to check the error.
func (q *ArrayBlockingQueue[T]) TryPop(ctx context.Context) (T, error) {
	// Acquire a slot in the semaphore, blocking if necessary.
	if err := q.notEmpty.Acquire(ctx, 1); err != nil {
		var zero T
		return zero, err
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	// Check if the context has already been cancelled when the lock is acquired.
	if ctx.Err() != nil {
		// Release a slot for the notEmpty semaphore.
		q.notEmpty.Release(1)
		var zero T
		return zero, ctx.Err()
	}

	t := q.items[q.head]
	q.head = (q.head + 1) % q.cap
	q.size--

	// Release a slot for the notFull semaphore.
	q.notFull.Release(1)

	return t, nil
}

// Size returns the number of elements in the queue, at the time of calling.
func (q *ArrayBlockingQueue[T]) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size
}

// Cap returns the capacity of the queue.
func (q *ArrayBlockingQueue[T]) Cap() int {
	return q.cap
}
