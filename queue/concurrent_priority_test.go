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
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/mizumoto-cn/fpkit/functional"
	xerr "github.com/mizumoto-cn/fpkit/internal/err"
	"github.com/mizumoto-cn/fpkit/queue"

	"github.com/stretchr/testify/assert"
)

type objCpq struct {
	Name string
	Age  int
}

var cmpCpq functional.ComparatorAny[objCpq] = func(a, b objCpq) bool {
	return a.Age < b.Age
}

func TestConcurrentPriorityQueue(t *testing.T) {
	cpq, err := queue.NewConcurrentPriorityQueue(cmpCpq, 3)
	assert.NoError(t, err)
	assert.NotNil(t, cpq)

	// Test Push
	err = cpq.Push(objCpq{Name: "Alice", Age: 30})
	assert.NoError(t, err)
	err = cpq.Push(objCpq{Name: "Bob", Age: 25})
	assert.NoError(t, err)
	err = cpq.Push(objCpq{Name: "Charlie", Age: 35})
	assert.NoError(t, err)

	// Test Size
	assert.Equal(t, 3, cpq.Size())

	// Test Full
	assert.True(t, cpq.Full())

	// Test Front
	front, err := cpq.Front()
	assert.NoError(t, err)
	assert.Equal(t, objCpq{Name: "Bob", Age: 25}, front)

	// Test Back
	back, err := cpq.Back()
	assert.NoError(t, err)
	assert.Equal(t, objCpq{Name: "Charlie", Age: 35}, back)

	// Test Pop
	popped, err := cpq.Pop()
	assert.NoError(t, err)
	assert.Equal(t, objCpq{Name: "Bob", Age: 25}, popped)

	// Test Size after Pop
	assert.Equal(t, 2, cpq.Size())

	// Test Empty
	assert.False(t, cpq.Empty())

	// Test Clear
	err = cpq.Clear()
	assert.NoError(t, err)
	assert.True(t, cpq.Empty())
}

// TODO: Fix the test
func TestConcurrentPriorityQueueRacing(t *testing.T) {
	cpq, err := queue.NewConcurrentPriorityQueue(cmpCpq, 100)
	assert.NoError(t, err)
	assert.NotNil(t, cpq)

	var wg sync.WaitGroup
	numGoroutines := 10
	numOperations := 1000

	// Push elements concurrently
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				err := cpq.Push(objCpq{Name: fmt.Sprintf("Name%d-%d", id, j), Age: rand.Intn(100)})
				assert.NoError(t, err)
			}
		}(i)
	}

	// Pop elements concurrently
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				_, err := cpq.Pop()
				if err != nil && err != xerr.ErrEmptyQueue {
					assert.NoError(t, err)
				}
			}
		}()
	}

	wg.Wait()

	// Check final size
	finalSize := cpq.Size()
	assert.True(t, finalSize >= 0 && finalSize <= 100)
}
