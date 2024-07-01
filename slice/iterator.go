package slice

import "github.com/mizumoto-cn/fpkit/internal/slice"

type Iterator[T any] struct {
	src    []T
	cursor int
}

func NewIterator[T any](src []T) *Iterator[T] {
	return &Iterator[T]{src: src}
}

func (i *Iterator[T]) HasNext() bool {
	return i.cursor < len(i.src)
}

func (i *Iterator[T]) Next() T {
	v := i.src[i.cursor]
	i.cursor++
	return v
}

func (i *Iterator[T]) Remove() {
	i.src, _, _ = slice.Delete(i.src, i.cursor-1)
}

func (i *Iterator[T]) Slice() []T {
	return i.src
}

func (i *Iterator[T]) Head() T {
	return i.src[0]
}

func (i *Iterator[T]) Tail() []T {
	return i.src[1:]
}

func (i *Iterator[T]) Last() T {
	return i.src[len(i.src)-1]
}

func (i *Iterator[T]) Init() []T {
	return i.src[:len(i.src)-1]
}

func (i *Iterator[T]) Index() int {
	return i.cursor - 1
}

func (i *Iterator[T]) Reset() {
	i.cursor = 0
}
