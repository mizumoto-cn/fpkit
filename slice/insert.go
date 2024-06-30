package slice

import (
	"github.com/mizumoto-cn/gogenerices/internal/slice"
)

func Insert[T any](s []T, index int, value T) (result []T, err error) {
	if index < 0 || index > len(s) {
		return nil, slice.ErrIndexOutOfRange
	}
	result = append(s[:index], append([]T{value}, s[index:]...)...)
	return
}
