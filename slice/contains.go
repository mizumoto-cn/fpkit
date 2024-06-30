package slice

import "github.com/mizumoto-cn/fpkit/functional"

// Contains returns true if the slice contains the value.
func Contains[T comparable](s []T, v T) bool {
	return ContainsFunc(s, func(x T) bool { return x == v })
}

// ContainsFunc returns true if the slice contains an element that satisfies the predicate.
func ContainsFunc[T any](s []T, eq func(T) bool) bool {
	for _, v := range s {
		if eq(v) {
			return true
		}
	}
	return false
}

// ContainsAny returns true if the slice contains any of the values.
func ContainsAny[T comparable](s []T, vs ...T) bool {
	for _, v := range vs {
		if Contains(s, v) {
			return true
		}
	}
	return false
}

// ContainsAll checks if all elements of subset are in the main slice.
func ContainsAll[T comparable](mainSlice, subset []T) bool {
	// Use a map to track occurrences in the main slice for quick lookup
	elementCount := make(map[T]int)
	for _, elem := range mainSlice {
		elementCount[elem]++
	}

	// Use Reduce to check if all elements of the subset are in the main slice
	return functional.Reduce(subset, func(acc bool, elem T) bool {
		if acc && elementCount[elem] > 0 {
			elementCount[elem]--
			return true
		}
		return false
	}, true)
}
