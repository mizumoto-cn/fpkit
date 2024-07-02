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
package slice_test

import (
	"testing"

	"github.com/mizumoto-cn/fpkit/slice"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	src := []int{1, 2, 3, 4}
	it := slice.NewIterator(src)

	//  HasNext  Next
	for i := 0; i < len(src); i++ {
		assert.True(t, it.HasNext())
		assert.Equal(t, src[i], it.Next())
	}

	// There should be no more elements
	assert.False(t, it.HasNext())

	// Reset the iterator
	it.Reset()
	assert.True(t, it.HasNext())

	// Head, Tail, Last, Init
	assert.Equal(t, src[0], it.Head())
	assert.Equal(t, src[1:], it.Tail())
	assert.Equal(t, src[len(src)-1], it.Last())
	assert.Equal(t, src[:len(src)-1], it.Init())

	// Index Remove
	it.Reset()
	it.Next()
	assert.Equal(t, 0, it.Index())
	it.Remove()
	assert.Equal(t, []int{2, 3, 4}, it.Slice())

	// Slice
	it.Reset()
	assert.Equal(t, []int{2, 3, 4}, it.Slice())
}
