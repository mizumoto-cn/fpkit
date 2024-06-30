package slice_test

import (
	"testing"

	"github.com/mizumoto-cn/fpkit/slice"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	cases := []struct {
		title string
		src   []int
		value int
		want  bool
	}{
		{
			title: "empty",
			src:   []int{},
			value: 1,
			want:  false,
		},
		{
			title: "single",
			src:   []int{1},
			value: 1,
			want:  true,
		},
		{
			title: "multiple",
			src:   []int{1, 2, 3, 4, 5},
			value: 3,
			want:  true,
		},
		{
			title: "negative",
			src:   []int{-1, -2, -3, -4, -5},
			value: -3,
			want:  true,
		},
		{
			title: "mixed",
			src:   []int{-1, 2, -3, 4, -5},
			value: 4,
			want:  true,
		},
		{
			title: "not found",
			src:   []int{1, 2, 3, 4, 5},
			value: 6,
			want:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.want, slice.Contains(c.src, c.value))
		})
	}
}
