package functional_test

import (
	"testing"

	"github.com/mizumoto-cn/fpkit/functional"

	assert "github.com/stretchr/testify/assert"
)

func TestFoldl(t *testing.T) {
	cases := []struct {
		title string
		src   []int
		want  int
		fn    func(int, int) int
		init  int
	}{
		{
			title: "Sum of integers",
			src:   []int{1, 2, 3, 4},
			want:  10,
			fn:    func(a, b int) int { return a + b },
			init:  0,
		},
		{
			title: "Product of integers",
			src:   []int{1, 2, 3, 4},
			want:  24,
			fn:    func(a, b int) int { return a * b },
			init:  1,
		},
		{
			title: "From left to right",
			src:   []int{9, 2, 1},
			want:  1, // 11 - 9 - 1
			fn: func(a, b int) int {
				if a > b {
					return a - b
				}
				return a
			},
			init: 11,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.want, functional.Foldl(c.src, c.fn, c.init))
		})
	}

}

func TestFoldr(t *testing.T) {
	cases := []struct {
		title string
		src   []int
		want  int
		fn    func(int, int) int
		init  int
	}{
		{
			title: "Sum of integers",
			src:   []int{1, 2, 3, 4},
			want:  10,
			fn:    func(a, b int) int { return a + b },
			init:  0,
		},
		{
			title: "Product of integers",
			src:   []int{1, 2, 3, 4},
			want:  24,
			fn:    func(a, b int) int { return a * b },
			init:  1,
		},
		{
			title: "From left to right",
			src:   []int{9, 2, 1},
			want:  8, // 11 - 1 - 2
			fn: func(a, b int) int {
				if a > b {
					return a - b
				}
				return a
			},
			init: 11,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.want, functional.Foldr(c.src, c.fn, c.init))
		})
	}
}
