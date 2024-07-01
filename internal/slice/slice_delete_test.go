package slice_test

import (
	"testing"

	"github.com/mizumoto-cn/fpkit/internal/slice"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	t.Run("Delete an element from the slice", func(t *testing.T) {
		t.Parallel()

		s := []int{1, 2, 3, 4, 5}
		s, deleted, err := slice.Delete(s, 2)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 4, 5}, s)
		assert.Equal(t, 3, deleted)
	})

	t.Run("Delete an element from the slice with the first index", func(t *testing.T) {
		t.Parallel()

		s := []int{1, 2, 3, 4, 5}
		s, deleted, err := slice.Delete(s, 0)
		assert.NoError(t, err)
		assert.Equal(t, []int{2, 3, 4, 5}, s)
		assert.Equal(t, 1, deleted)
	})

	t.Run("Delete an element from the slice with the last index", func(t *testing.T) {
		t.Parallel()

		s := []int{1, 2, 3, 4, 5}
		s, deleted, err := slice.Delete(s, 4)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4}, s)
		assert.Equal(t, 5, deleted)
	})

	t.Run("Delete an element from the slice with an invalid index", func(t *testing.T) {
		t.Parallel()

		s := []int{1, 2, 3, 4, 5}
		s, deleted, err := slice.Delete(s, 5)
		assert.Error(t, err)
		assert.Nil(t, s)
		assert.Zero(t, deleted)
	})
}
