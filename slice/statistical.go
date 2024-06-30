package slice

import (
	"github.com/mizumoto-cn/fpkit/functional"
)

// Max returns the maximum value in the Orderable slice.
// If the slice is empty, it panics.
// Which was also the case in the original Go "slices" library.
func Max[T functional.Orderable](s []T) T {
	res := s[0]
	for _, v := range s[1:] {
		if res < v {
			res = v
		}
	}
	return res
}

// Min returns the minimum value in the Orderable slice.
// If the slice is empty, it returns a NewIndexOutOfRangeError.
func Min[T functional.Orderable](s []T) T {
	res := s[0]
	for _, v := range s[1:] {
		if res > v {
			res = v
		}
	}
	return res
}

// You can also use this to find out the extreme value by changing the comparison function.
// For non-Orderable types, you can also use the following function.
// If the slice is empty, it panics.
//
// cmp: if you want max, make it return `true` when its lhs > rhs.
func ExtremeValue[T any](s []T, cmp func(T, T) bool) T {
	res := s[0]
	for _, v := range s[1:] {
		if !cmp(res, v) {
			res = v
		}
	}
	return res
}

// Sum returns the sum of the slice.
// If the slice is empty, it returns zero value.
func Sum[T functional.Numeric](s []T) T {
	var res T
	for _, v := range s {
		res += v
	}
	return res
}

// Roadmap:
// 1. Add variance/standard deviation/median/quartile functions. (v1.0.1~)
