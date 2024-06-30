package slice_test

import (
	"testing"

	"github.com/mizumoto-cn/fpkit/internal/err"
	"github.com/mizumoto-cn/fpkit/slice"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		title       string
		slice       []int
		index       int
		value       int
		want        []int
		expectedErr error
	}{
		{
			title: "insert to the head",
			slice: []int{1, 2, 3},
			index: 0,
			value: 0,
			want:  []int{0, 1, 2, 3},
		},
		{
			title: "insert to the middle",
			slice: []int{1, 2, 3},
			index: 1,
			value: 4,
			want:  []int{1, 4, 2, 3},
		},
		{
			title: "insert to the tail",
			slice: []int{1, 2, 3},
			index: 3,
			value: 4,
			want:  []int{1, 2, 3, 4},
		},
		{
			title:       "insert to the out of range",
			slice:       []int{1, 2, 3},
			index:       4,
			value:       4,
			expectedErr: err.NewIndexOutOfRangeError(4, 3),
		},
		{
			title:       "insert to the negative index",
			slice:       []int{1, 2, 3},
			index:       -1,
			value:       4,
			expectedErr: err.NewIndexOutOfRangeError(-1, 3),
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got, err := slice.Insert(c.slice, c.index, c.value)
			if c.expectedErr != nil {
				assert.Equal(t, c.expectedErr, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}
