package functional

import (
	"sort"
)

type FnObject func(any) any

var _ = FnObject(nil)

type Real interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Numeric interface {
	Real | ~complex64 | ~complex128
}

func Sum[T Numeric](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

type Orderable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

type Comparator[T Orderable] func(T, T) bool

func CompareTo[T Orderable](a, b T) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func Sort[T Orderable](a []T, cmp Comparator[T]) {
	sort.SliceStable(a, func(i, j int) bool {
		return cmp(a[i], a[j])
	})
}

func SortOrdered[T Orderable](ascending bool, input ...T) []T {
	if ascending {
		Sort(input, func(a, b T) bool {
			return CompareTo(a, b) < 0
			// a < b, ascending
		})
	} else {
		Sort(input, func(a, b T) bool {
			return CompareTo(a, b) > 0
			// a > b, descending
		})
	}
	return input
}

func SortAsc[T Orderable](input ...T) []T {
	return SortOrdered(true, input...)
}

func SortDesc[T Orderable](input ...T) []T {
	return SortOrdered(false, input...)
}
