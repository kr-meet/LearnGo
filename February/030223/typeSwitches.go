package main

import (
	"fmt"
)

func f(inter interface{}) {

	switch t := inter.(type) {

	case int:
		fmt.Println(t)

	case string:
		fmt.Println(t)

	default:
		fmt.Println("Unknown")
	}
}

func main() {

	f("string")
	f(654)
	f(433.232)
}
