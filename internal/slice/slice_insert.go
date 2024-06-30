package slice

import (
	"github.com/mizumoto-cn/fpkit/internal/err"
)

func Insert[T any](s []T, index int, value T) ([]T, error) {
	if index < 0 || index > len(s) {
		return nil, err.NewIndexOutOfRangeError(index, len(s))
	}
	var zeroValue T
	s = append(s, zeroValue)
	// `copy` operates the memory directly, so it is faster than a for loop
	copy(s[index+1:], s[index:])
	s[index] = value
	return s, nil
}
