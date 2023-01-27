package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	isOkay bool       = false
	point  float64    = 2.31
	z      complex128 = cmplx.Sqrt(-1 + 12i)
)

const (
	big   = 1 << 100
	small = big >> 99
)

func getFloat(x float64) float64 {
	return x * 0.1
}

func pow(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	} else {
		fmt.Println("v >= lim", "%g" >= "%g", v, lim)
		return lim
	}
}

func infinite() {
	for {
	}
}

type Vertex struct {
	X, Y int
}

func main() {
	// fmt.Printf("You have %g problems. \n", math.Pi)

	// fmt.Println(add(5, 4))

	// a, b := swap("Knack", "Root")
	// fmt.Println(a, b)

	// fmt.Println(split(17))

	// var i int
	// var c, java,' python = true, false, "kfnsdknf!"
	// fmt.Println(i, c, java, python)

	// var i, j = 1, 10
	// k := "String"
	// fmt.Println(i, j, k)

	// fmt.Printf("Type is %T and value is %v \n", isOkay, isOkay)
	// fmt.Printf("Type is %T and value is %v \n", point, point)
	// fmt.Printf("Type is %T and value is %v \n", z, z)

	// var i, j, k = 3, 3.14, 6.435341
	// var newI = float64(i)
	// newI += 1.22
	// var newJ = int(j)
	// var newK = int(k)
	// fmt.Println(newI, newJ, newK)

	// const value = 2312
	// fmt.Println(value)
	// var f float64 = 10000000000000000000000045
	// fmt.Println(getFloat(big))
	// fmt.Println(f)

	// sum := 0
	// for i := 0; i <= 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// sum := 1
	// for sum <= 1000;{
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// for {
	// 	fmt.Println("here")
	// }
	// fmt.Println("Not here")

	// fmt.Println(pow(3, 3, 10))

	// var os = "linux"
	// switch os {
	// case "darwin":
	// 	fmt.Println("DARWIN")
	// fmt.Println("2")
	// case "linux":
	// 	fmt.Println("LINUX")
	// 	fmt.Println(22)
	// default:
	// 	fmt.Println("DEFAULT")
	// }

	// day := time.Now().Weekday()
	// switch time.Saturday {
	// case day + 0:
	// 	fmt.Println("Today")
	// case day + 1:
	// 	fmt.Println("Tomorrow")
	// case day + 2:
	// 	fmt.Println("2 days left")
	// default:
	// 	fmt.Println("Time there")
	// }

	// defer fmt.Println("World")
	// fmt.Println("Hello")

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("Lastline")

	// infinite()
	// defer fmt.Println("here")

	// var i, j = 20, 30
	// p := &i
	// fmt.Println(&i)
	// fmt.Println(p, *p)
	// *p = 20
	// fmt.Println(i, p)
	// p = &j
	// fmt.Println(p, *p)

	// vertex := Vertex{1, 3, "KnackRoot"}
	// fmt.Println(vertex)
	// vertex.X = 3
	// fmt.Println(vertex)

	// vertex := Vertex{10, 20, "Hey"}
	// p := &vertex
	// p.X = 100
	// vertex.Y = 100
	// fmt.Println(vertex)

	// fmt.Println(v1, v2, v3, p)

	v1 := Vertex{1, 2}
	fmt.Println(v1)
}
