package slice

// Union of two slices, removing duplicates.
func Union[T comparable](a, b []T) []T {
	set := make(map[T]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}
	for _, v := range b {
		set[v] = struct{}{}
	}
	result := make([]T, 0, len(set))
	for v := range set {
		result = append(result, v)
	}
	return result
}

// Intersection of two slices.
func Intersection[T comparable](a, b []T) []T {
	set := make(map[T]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}
	result := make([]T, 0, len(set))
	for _, v := range b {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
