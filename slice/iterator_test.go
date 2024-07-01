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
