package slice

// Index returns the index of the first occurrence of the value v in the slice src, or -1 if not present.
func Index[T comparable](src []T, v T) int {
	return IndexMatchFunc(src, func(e T) bool { return e == v })
}

// IndexMatchFunc returns the index of the first element in the slice that satisfies the provided function, or -1 if not present.
func IndexMatchFunc[T any](src []T, match func(T) bool) int {
	for k, v := range src {
		if match(v) {
			return k
		}
	}
	return -1
}

// IndexAllMatchFunc returns the index of the all elements in the slice that satisfies the provided function.
func IndexAllMatchFunc[T any](src []T, match func(T) bool) []int {
	var indexes []int
	for k, v := range src {
		if match(v) {
			indexes = append(indexes, k)
		}
	}
	return indexes
}

// IndexAll returns the index of the all elements in the slice that satisfies the provided function.
func IndexAll[T comparable](src []T, v T) []int {
	return IndexAllMatchFunc(src, func(e T) bool { return e == v })
}

// LastIndexMatchFunc returns the index of the last element in the slice that satisfies the provided function, or -1 if not present.
func LastIndexMatchFunc[T any](src []T, match func(T) bool) int {
	for i := len(src) - 1; i >= 0; i-- {
		if match(src[i]) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last occurrence of the value v in the slice src, or -1 if not present.
func LastIndex[T comparable](src []T, v T) int {
	return LastIndexMatchFunc(src, func(e T) bool { return e == v })
}
