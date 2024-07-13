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

	"github.com/mizumoto-cn/fpkit/queue"

	"github.com/stretchr/testify/assert"
)

func TestBasicQueue(t *testing.T) {
	q := *queue.NewBasicQueue[int](10)

	var _ queue.Queue[int] = &q

	assert.NotNil(t, q)

	assert.True(t, q.Empty())

	assert.Equal(t, 10, q.Cap())

	for i := 0; i < 10; i++ {
		assert.NoError(t, q.Push(i))
	}

	assert.False(t, q.Empty())

	assert.Equal(t, 10, q.Size())

	for i := 0; i < 10; i++ {
		tail, err := q.Back()
		assert.Equal(t, 9, tail)
		assert.NoError(t, err)
		v, err := q.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
		assert.Equal(t, 10-i-1, q.Size())
	}

	assert.True(t, q.Empty())

	assert.Equal(t, 10, q.Cap())

	assert.Equal(t, 0, q.Size())

	assert.NoError(t, q.Push(1))

	assert.NoError(t, q.Clear())

	assert.True(t, q.Empty())

	assert.Equal(t, 10, q.Cap())

	assert.Equal(t, 0, q.Size())

	_, err := q.Pop()

	assert.Error(t, err)

	_, err = q.Front()

	assert.Error(t, err)

	_, err = q.Back()

	assert.Error(t, err)

	for i := 0; i < 10; i++ {
		assert.NoError(t, q.Push(i))
	}

	assert.NoError(t, q.Resize(5, true))

	assert.Equal(t, 5, q.Size())

	assert.Equal(t, 5, q.Cap())

	assert.Error(t, q.Push(1))

	front, err := q.Front()
	assert.NoError(t, err)
	assert.Equal(t, 0, front)

	for i := 0; i < 5; i++ {
		v, err := q.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
	}
	// print(q.Size())

	// test resize with shrink false

	for i := 0; i < 5; i++ {
		assert.NoError(t, q.Push(i))
	}

	assert.Error(t, q.Resize(3, false))

	assert.Equal(t, 5, q.Size())

	assert.Equal(t, 5, q.Cap())

	assert.NoError(t, q.Resize(10, false))

	assert.Equal(t, 5, q.Size())

	assert.Equal(t, 10, q.Cap())

	for i := 0; i < 5; i++ {
		assert.NoError(t, q.Push(i+5))
	}

	assert.Error(t, q.Resize(-1, true))

	q2 := *queue.NewBasicQueue[int](10)
	assert.NoError(t, q2.Push(1))
	assert.NoError(t, q2.Push(2))

	// Test Swap
	q.Swap(&q2)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 10, q2.Size())
	tmp, err := q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, tmp)
	tmp, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, tmp)
	for i := 0; i < 10; i++ {
		tmp, err = q2.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, tmp)
	}
}
