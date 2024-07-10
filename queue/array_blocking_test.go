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

	err := q.Push(context.Background(), 1)
	assert.NoError(t, err)
	println("1 pushed")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	println("context created %v", ctx)
	defer cancel()
	err = q.Push(ctx, 1)
	println("2 tried to push: %v", err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)

	// TryPop the element to make space in the queue
	elem, err := q.TryPop(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, elem)

	// Now we should be able to Push again without error
	err = q.Push(context.Background(), 2)
	assert.NoError(t, err)
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
