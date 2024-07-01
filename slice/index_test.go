package slice_test

import (
	"testing"

	"github.com/mizumoto-cn/fpkit/slice"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	cases := []struct {
		title    string
		src      []int
		value    int
		expected int
	}{
		{
			title:    "value present at beginning",
			src:      []int{1, 2, 3, 4},
			value:    1,
			expected: 0,
		},
		{
			title:    "value present in middle",
			src:      []int{1, 2, 3, 4},
			value:    3,
			expected: 2,
		},
		{
			title:    "value not present",
			src:      []int{1, 2, 3, 4},
			value:    5,
			expected: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.Index(c.src, c.value)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestIndexAll(t *testing.T) {
	var empty []int
	cases := []struct {
		title    string
		src      []int
		value    int
		expected []int
	}{
		{
			title:    "multiple occurrences",
			src:      []int{1, 2, 3, 1, 1},
			value:    1,
			expected: []int{0, 3, 4},
		},
		{
			title:    "no occurrences",
			src:      []int{1, 2, 3, 4},
			value:    5,
			expected: empty,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.IndexAll(c.src, c.value)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestLastIndex(t *testing.T) {
	cases := []struct {
		title    string
		src      []int
		value    int
		expected int
	}{
		{
			title:    "value present at end",
			src:      []int{1, 2, 3, 4},
			value:    4,
			expected: 3,
		},
		{
			title:    "value present in middle",
			src:      []int{1, 4, 3, 4},
			value:    4,
			expected: 3,
		},
		{
			title:    "value not present",
			src:      []int{1, 2, 3, 4},
			value:    5,
			expected: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.LastIndex(c.src, c.value)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestIndexMatchFunc(t *testing.T) {
	cases := []struct {
		title    string
		src      []int
		match    func(int) bool
		expected int
	}{
		{
			title:    "match found",
			src:      []int{1, 2, 3, 4},
			match:    func(i int) bool { return i == 3 },
			expected: 2,
		},
		{
			title:    "match not found",
			src:      []int{1, 2, 3, 4},
			match:    func(i int) bool { return i == 5 },
			expected: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.IndexMatchFunc(c.src, c.match)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestIndexAllMatchFunc(t *testing.T) {
	var empty []int
	cases := []struct {
		title    string
		src      []int
		match    func(int) bool
		expected []int
	}{
		{
			title:    "multiple matches",
			src:      []int{1, 2, 3, 1, 1},
			match:    func(i int) bool { return i == 1 },
			expected: []int{0, 3, 4},
		},
		{
			title:    "no matches",
			src:      []int{1, 2, 3, 4},
			match:    func(i int) bool { return i == 5 },
			expected: empty,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			result := slice.IndexAllMatchFunc(c.src, c.match)
			assert.Equal(t, c.expected, result)
		})
	}
}
