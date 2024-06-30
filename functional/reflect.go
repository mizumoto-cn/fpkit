package functional

import "reflect"

func PtrOf[T any](v T) *T {
	return &v
}

func SliceOf[T any](v ...T) []T {
	return v
}

func IsPtr[T any](v T) bool {
	return Kind(v) == reflect.Ptr
}

func Kind[T any](v T) reflect.Kind {
	return reflect.TypeOf(v).Kind()
}

func IsNil[T any](v T) bool {
	// 2 cases, pointer or not
	val := reflect.ValueOf(v)
	if Kind(v) == reflect.Ptr {
		return val.IsNil()
	}
	return !val.IsValid()
}
