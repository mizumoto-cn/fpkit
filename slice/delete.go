package slice

import (
	"github.com/mizumoto-cn/fpkit/internal/slice"
)

// Delete removes the element at the given index from the slice.
func Delete[T any](src []T, index int) ([]T, error) {
	s, _, err := slice.Delete(src, index)
	return s, err
}

// DeleteMatched removes all elements that matches the given function from the slice.
func DeleteMatched[T any](src []T, match func(T) bool) []T {
	pos := 0
	for idx := range src {
		if match(src[idx]) {
			continue
		}
		src[pos] = src[idx]
		pos++
	}
	return src[:pos]
}
