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
package queue_test

import (
	"context"
	"testing"
	"time"

	"github.com/mizumoto-cn/fpkit/queue"

	"github.com/stretchr/testify/assert"
)

func TestArrayBlockingQueue1(t *testing.T) {
	q := queue.NewArrayBlockingQueue[int](1)
	assert.NotNil(t, q)
	println("q created")
	assert.Zero(t, q.Size())
	assert.Equal(t, 1, q.Cap())

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	out, err := q.TryPop(ctx)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
	assert.Zero(t, out)

	err = q.Push(context.Background(), 1)
	assert.NoError(t, err)
	println("1 pushed")

	ctx1, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	println("context created ", ctx1)
	defer cancel()
	err = q.Push(ctx1, 1)
	println("2 tried to push: ", err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)

	// TryPop the element to make space in the queue
	elem, err := q.TryPop(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, elem)

	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	elem, err = q.TryPop(ctx2)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
	assert.Zero(t, elem)

	// Now we should be able to Push again without error
	err = q.Push(context.Background(), 2)
	assert.NoError(t, err)

	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 1, q.Cap())
}

// TestArrayBlockingQueue2 tests the blocking behavior of ArrayBlockingQueue with multiple goroutines.
func TestArrayBlockingQueue2(t *testing.T) {
	q := queue.NewArrayBlockingQueue[int](1)
	assert.NotNil(t, q)
	println("q created")

	// Push an element to the queue
	err := q.Push(context.Background(), 1)
	assert.NoError(t, err)
	println("1 pushed")

	// Create a channel to synchronize the goroutines
	ch := make(chan struct{})

	// Create a goroutine to push an element to the queue
	go func() {
		err := q.Push(context.Background(), 2)
		assert.NoError(t, err)
		println("2 pushed")
		ch <- struct{}{}
	}()

	// Create a goroutine to pop an element from the queue
	go func() {
		elem, err := q.TryPop(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 1, elem)
		println("1 popped")
		ch <- struct{}{}
	}()

	// Wait for the goroutines to finish
	<-ch
	<-ch
}

func TestArrayBlockingQueueRacePush(t *testing.T) {
	q := queue.NewArrayBlockingQueue[int](10000000)
	for i := range 10000000 {
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
			defer cancel()
			if err := q.Push(ctx, i); err != nil {
				println("push error: ", err.Error())
				// "push error:  context deadline exceeded" for sure
			}
			// q.Push(ctx, i)
		}()
	}
	<-time.After(10 * time.Millisecond)
	print("10000000 threads pushed with 10ms timeout, the result is:", q.Size())
	// things like 10000000 threads pushed with 10ms timeout, the result is: 19125
	assert.Less(t, q.Size(), 10000000)
	// usually the result is less than 10000000, like 19125
}

func TestArrayBlockingQueueRacePop(t *testing.T) {
	q := queue.NewArrayBlockingQueue[int](10000000)
	for i := range 10000000 {
		err := q.Push(context.Background(), i)
		assert.NoError(t, err)
	}
	for range 10000000 {
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
			defer cancel()
			if _, err := q.TryPop(ctx); err != nil {
				println("pop error: ", err.Error())
				// "pop error:  context deadline exceeded" for sure
			}
			// q.TryPop(ctx)
		}()
	}
	<-time.After(100 * time.Millisecond)
	print("10000000 threads popped with 10ms timeout, the result is:", q.Size())
	// things like 10000000 threads popped with 10ms timeout, the result is: 875
	assert.Greater(t, q.Size(), 0)
	// usually the result is greater than 0, like 675
}
