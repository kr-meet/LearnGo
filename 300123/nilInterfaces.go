package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("Nil it is")
		return
	}
	fmt.Println(t.S)

}

func show(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	var t *T
	i = t
	show(i)
	i.M()

	i = &T{"hello"}
	show(i)
	i.M()
}
