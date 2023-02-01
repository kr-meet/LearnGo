package main

import "fmt"

type Shape interface {
	Area() float64
	Volume() float64
}

type Circle struct {
	R float64
}

type Square struct {
	L float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.R * c.R
}

func (c Circle) Volume() float64 {
	return 4 / 3 * 3.14 * c.R * c.R * c.R
}

func (s Square) Area() float64 {
	return s.L * s.L
}

func (s Square) Volume() float64 {
	return s.L * s.L * s.L
}

func main() {

	var a Shape
	c1 := Circle{4}
	a = c1
	fmt.Println(a.Area(), a.Volume())

	s1 := Square{4}
	a = s1
	fmt.Println(a.Area(), a.Volume())
}
