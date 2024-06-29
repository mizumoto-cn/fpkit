package slice_test

import (
	"testing"

	"github.com/mizumoto-cn/gogenerics/internal/err"
	"github.com/mizumoto-cn/gogenerics/internal/slice"

	assert "github.com/stretchr/testify/assert"
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
			if c.expectedErr != nil {
				assert.PanicsWithError(t, c.expectedErr.Error(), func() {
					slice.Insert(c.slice, c.index, c.value)
				})
			} else {
				assert.Equal(t, c.want, slice.Insert(c.slice, c.index, c.value))
			}
		})
	}
}
