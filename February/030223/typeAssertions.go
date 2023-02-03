package main

import "fmt"

func main() {

	var i interface{} = "Hey There!"
	fmt.Printf("%T", i)
	fmt.Println()
	fmt.Println(i)

	s, ok := i.(float64)
	fmt.Println(s, ok)
	fmt.Println(i.(int))
}
