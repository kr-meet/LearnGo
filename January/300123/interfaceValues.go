package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

type Float float64

func (t *T) M() {
	fmt.Println(t.S)
}

func (f Float) M() {
	fmt.Println(f)
}

func show(i I) {
	fmt.Printf("%v, %T", i, i)
	fmt.Println()
}

func main() {

	var i I

	i = &T{"hello"}
	show(i)
	i.M()

	i = Float(3.14)
	show(i)
	i.M()
}
