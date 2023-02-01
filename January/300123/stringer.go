package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {

	p1 := Person{"Meet", 21}
	p2 := Person{"Jarvis", 21}
	p3 := Person{"Jarvis", 201}
	fmt.Println(p3, p2)
	fmt.Println(p1, p2)
}
