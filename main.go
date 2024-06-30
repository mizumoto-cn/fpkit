package main

import (
	"slices"

	"github.com/mizumoto-cn/fpkit/functional"
)

func main() {

	println(functional.Sum(1, 2, 3+4i))

	mlist := []int{}

	m := slices.Max[[]int](mlist)
	print(m)
}
