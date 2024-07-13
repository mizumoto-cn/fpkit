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
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/mizumoto-cn/fpkit/functional"
	"github.com/mizumoto-cn/fpkit/internal/err"
	"github.com/mizumoto-cn/fpkit/queue"

	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]
		val  int

		wanted []int
		wErr   error
	}{
		{
			name: "push to empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			val:    1,
			wanted: []int{1},
		},
		{
			name: "push to non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				return q
			},
			val:    2,
			wanted: []int{1, 2},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			err := q.Push(tt.val)
			require.Equal(t, tt.wErr, err)
			require.Equal(t, tt.wanted, functional.SortAsc(q.Slice()...))
		})
	}
}

func TestPop(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]

		wanted int
		wErr   error
	}{
		{
			name: "pop from empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			wanted: 0,
			wErr:   err.ErrEmptyQueue,
		},
		{
			name: "pop from non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				return q
			},
			wanted: 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			val, err := q.Pop()
			require.Equal(t, tt.wErr, err)
			require.Equal(t, tt.wanted, val)
		})
	}
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]

		wanted bool
	}{
		{
			name: "empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			wanted: true,
		},
		{
			name: "non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				return q
			},
			wanted: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			require.Equal(t, tt.wanted, q.Empty())
		})
	}
}

func TestSize(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]

		wanted int
	}{
		{
			name: "empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			wanted: 0,
		},
		{
			name: "non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				return q
			},
			wanted: 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			require.Equal(t, tt.wanted, q.Size())
		})
	}
}

func TestCap(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]

		wanted int
	}{
		{
			name: "empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			wanted: -1,
		},
		{
			name: "non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				return q
			},
			wanted: -1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			require.Equal(t, tt.wanted, q.Cap())
		})
	}
}

func TestBack(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]

		wanted int
		wErr   error
	}{
		{
			name: "back from empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			wanted: 0,
			wErr:   err.ErrEmptyQueue,
		},
		{
			name: "back from non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				return q
			},
			wanted: 1,
		},
		{
			name: "back from non-empty queue with multiple elements",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				err = q.Push(2)
				require.NoError(t, err)
				return q
			},
			wanted: 2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			val, err := q.Back()
			require.Equal(t, tt.wErr, err)
			require.Equal(t, tt.wanted, val)
		})
	}
}

func TestFront(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]

		wanted int
		wErr   error
	}{
		{
			name: "front from empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			wanted: 0,
			wErr:   err.ErrEmptyQueue,
		},
		{
			name: "front from non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				return q
			},
			wanted: 1,
		},
		{
			name: "front from non-empty queue with multiple elements",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				err = q.Push(2)
				require.NoError(t, err)
				return q
			},
			wanted: 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			val, err := q.Front()
			require.Equal(t, tt.wErr, err)
			require.Equal(t, tt.wanted, val)
		})
	}
}

func TestClear(t *testing.T) {
	t.Parallel()
	var zeroSlice []int
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]

		wanted []int
	}{
		{
			name: "clear empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			wanted: zeroSlice,
		},
		{
			name: "clear non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				err = q.Push(2)
				require.NoError(t, err)
				return q
			},
			wanted: zeroSlice,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			q.Clear()
			require.Equal(t, tt.wanted, q.Slice())
		})
	}
}

func TestSlice(t *testing.T) {
	t.Parallel()
	var zeroSlice []int
	cases := []struct {
		name string
		q    func() *queue.LinkedQueue[int]

		wanted []int
	}{
		{
			name: "empty queue",
			q: func() *queue.LinkedQueue[int] {
				return queue.NewLinkedQueue[int]()
			},
			wanted: zeroSlice,
		},
		{
			name: "non-empty queue",
			q: func() *queue.LinkedQueue[int] {
				q := queue.NewLinkedQueue[int]()
				err := q.Push(1)
				require.NoError(t, err)
				err = q.Push(2)
				require.NoError(t, err)
				return q
			},
			wanted: []int{1, 2},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			require.Equal(t, tt.wanted, q.Slice())
		})
	}
}

func TestRacing(t *testing.T) {
	t.Parallel()
	q := queue.NewLinkedQueue[int]()
	var wg sync.WaitGroup
	wg.Add(100000)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				v := rand.Intn(1000)
				q.Push(j + v)
			}
		}()
	}

	var cnt int32
	for i := 0; i < 100; i++ {
		go func() {
			for {
				if atomic.LoadInt32(&cnt) >= 100000 {
					break
				}
				_, err := q.Pop()
				if err == nil {
					atomic.AddInt32(&cnt, 1)
					wg.Done()
				}
			}
		}()
	}
	wg.Wait()
	require.Equal(t, 0, q.Size())
}
