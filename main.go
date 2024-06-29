package main

import "github.com/mizumoto-cn/gogenerics/functional"

type outer struct {
	inner
}

func (o outer) Name() string {
	return "outer"
}

func (i inner) Name() string {
	return "inner"
}

type inner struct {
}

func (i inner) Say() {
	println("hello," + i.Name())
}

func main() {
	o := outer{}
	o.Say()

	println(functional.Sum(1, 2, 3+4i))

}
