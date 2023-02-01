package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y int
}

type data float64

func add(v Vertex) int {
	return v.X + v.Y
}

func (v Vertex) add() int {
	return v.X + v.Y
}

func (d1 data) print(d2 data) {
	fmt.Println(d1, d2)
}

func (d data) abs() float64 {

	if d < 0 {
		return float64(-1 * d)
	}

	return 0
}

func (v *Vertex) change(c int) {
	v.X += c
	v.Y += c
	fmt.Println(v)
}

func changeValue(v *Vertex, val int) {
	v.X += val
	v.Y += val
}

func (v *Vertex) scale(val int) {
	v.X += val
	v.Y += val
}

func scale(v *Vertex, val int) {
	v.X += val
	v.Y += val
}

func (v Vertex) plus() int {
	return v.X + v.Y
}

func plus(v Vertex) {
	fmt.Println(v.X + v.Y)
}

func main() {

	v := Vertex{1, 2}
	fmt.Println(v)
	d1 := data(2)
	d2 := data(3)
	d1.print(d2)
	fmt.Println(add(v))
	fmt.Println(v.add())

	d := data(math.Sqrt2)
	fmt.Println(d.abs())

	var x = 10
	y := &x
	*y = 100
	fmt.Println(x)

	v := Vertex{1, 2}
	v1 := &v
	v1.change(10)
	fmt.Println(v)
	fmt.Printf("%T", v1)

	v := Vertex{1, 2}
	v.change(20)
	fmt.Println(v)

	v := Vertex{10, 19}
	changeValue(&v, 10)
	fmt.Println(v)

	v := Vertex{1, 2}
	v.scale(10)
	fmt.Println(v)
	scale(&v, 100)
	fmt.Println(v)

	v := Vertex{10, 20}
	fmt.Println(v.plus())
	plus(v)
	p := &Vertex{1, 2}
	fmt.Println(p.plus())
	fmt.Println()
	plus(*p)
}
