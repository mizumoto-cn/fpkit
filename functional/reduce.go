package functional

// Foldl applies a function to each element of a list, starting from the left, and returns the final value.
func Foldl[T any, U any](s []T, fn func(U, T) U, init U) U {
	for _, v := range s {
		init = fn(init, v)
	}
	return init
}

// Foldr applies a function to each element of a list, starting from the right, and returns the final value.
func Foldr[T any, U any](s []T, fn func(U, T) U, init U) U {
	for i := len(s) - 1; i >= 0; i-- {
		init = fn(init, s[i])
	}
	return init
}

// Reduce applies a reduction function to each element of the slice on a left-to-right basis to a given initial value.
func Reduce[T any, U any](s []T, fn func(U, T) U, init U) U {
	// result := init
	// for _, v := range s {
	// 	result = fn(result, v)
	// }
	// return result
	return Foldl(s, fn, init)
}
