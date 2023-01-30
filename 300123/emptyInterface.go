package main

import "fmt"

func main() {
	var i interface {
	}

	i = 42
	show(i)
	i = "Knackroot"
	show(i)
}

func show(i interface{}) {
	fmt.Printf("%T,%v\n", i, i)
}
