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
	"testing"

	"github.com/mizumoto-cn/fpkit/functional"
	"github.com/mizumoto-cn/fpkit/queue"

	"github.com/stretchr/testify/assert"
)

type obj struct {
	Name string
	Age  int
}

var cmp functional.ComparatorAny[obj] = func(a, b obj) bool {
	return a.Age < b.Age
}

func TestPriorityQueue(t *testing.T) {
	pFail, err := queue.NewPriorityQueue[obj](cmp, 0)
	assert.Nil(t, pFail)
	assert.Error(t, err)

	pq, err := queue.NewPriorityQueue[obj](cmp, 3)
	assert.Nil(t, err)
	assert.NotNil(t, pq)
	assert.Zero(t, pq.Size())
	assert.Equal(t, 3, pq.Cap())
	front, err := pq.Front()
	assert.NotNil(t, err)
	assert.Equal(t, obj{Name: "", Age: 0}, front)
	back, err := pq.Back()
	assert.NotNil(t, err)
	assert.Equal(t, obj{Name: "", Age: 0}, back)
	isEmpty := pq.Empty()
	assert.True(t, isEmpty)
	isFull := pq.Full()
	assert.False(t, isFull)

	err = pq.Push(obj{Name: "Alice", Age: 20})
	assert.Nil(t, err)
	assert.Equal(t, 1, pq.Size())
	err = pq.Push(obj{Name: "Bob", Age: 30})
	assert.Equal(t, 2, pq.Size())
	assert.Nil(t, err)
	err = pq.Push(obj{Name: "Charlie", Age: 10})
	assert.Equal(t, 3, pq.Size())
	assert.Nil(t, err)
	isEmpty = pq.Empty()
	assert.False(t, isEmpty)
	isFull = pq.Full()
	assert.True(t, isFull)

	err = pq.Push(obj{Name: "David", Age: 40})
	assert.NotNil(t, err)
	assert.Equal(t, 3, pq.Size())

	v, err := pq.Pop()
	assert.Nil(t, err)
	assert.Equal(t, obj{Name: "Charlie", Age: 10}, v)
	assert.Equal(t, 2, pq.Size())

	v, err = pq.Pop()
	assert.Nil(t, err)
	assert.Equal(t, obj{Name: "Alice", Age: 20}, v)
	assert.Equal(t, 1, pq.Size())

	v, err = pq.Pop()
	assert.Nil(t, err)
	assert.Equal(t, obj{Name: "Bob", Age: 30}, v)
	assert.Zero(t, pq.Size())

	v, err = pq.Pop()
	assert.Equal(t, obj{Name: "", Age: 0}, v)
	assert.Zero(t, pq.Size())
	assert.NotNil(t, err)
	assert.Equal(t, 3, pq.Cap())

	err = pq.Push(obj{Name: "Alice", Age: 20})
	assert.Nil(t, err)
	assert.Equal(t, 1, pq.Size())
	err = pq.Push(obj{Name: "Bob", Age: 30})
	assert.Equal(t, 2, pq.Size())
	assert.Nil(t, err)
	err = pq.Push(obj{Name: "Charlie", Age: 10})
	assert.Equal(t, 3, pq.Size())
	assert.Nil(t, err)

	assert.NoError(t, pq.Clear())
	assert.Zero(t, pq.Size())
	assert.Equal(t, 3, pq.Cap())

	err = pq.Push(obj{Name: "Alice", Age: 20})
	assert.Nil(t, err)
	assert.Equal(t, 1, pq.Size())
	err = pq.Push(obj{Name: "Bob", Age: 30})
	assert.Equal(t, 2, pq.Size())
	assert.Nil(t, err)
	err = pq.Push(obj{Name: "Charlie", Age: 10})
	assert.Equal(t, 3, pq.Size())
	assert.Nil(t, err)

	v, err = pq.Front()
	assert.Nil(t, err)
	assert.Equal(t, obj{Name: "Charlie", Age: 10}, v)

	v, err = pq.Back()
	assert.Nil(t, err)
	assert.Equal(t, obj{Name: "Bob", Age: 30}, v)
}

// TODO: Add more test cases to cover all scenarios of `down` function in priority.go
