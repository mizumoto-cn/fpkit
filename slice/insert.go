package slice

import (
	"github.com/mizumoto-cn/gogenerics/internal/slice"
)

func Insert[T any](s []T, index int, value T) ([]T, error) {
	return slice.Insert(s, index, value)
}
