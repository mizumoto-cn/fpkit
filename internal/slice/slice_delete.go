package slice

import "github.com/mizumoto-cn/fpkit/internal/err"

func Delete[T any](src []T, index int) ([]T, T, error) {
	l := len(src)
	if index < 0 || index >= l {
		var zero T
		return nil, zero, err.NewIndexOutOfRangeError(index, l)
	}
	deleted := src[index]
	copy(src[index:], src[index+1:])
	return src[:l-1], deleted, nil
}
