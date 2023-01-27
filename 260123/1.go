package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func printSlice(str []string) {
	fmt.Printf("len = %d cap = %d %v\n", len(str), cap(str), str)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

var mp = map[string]Vertex{
	"Java":   Vertex{15, 20},
	"Python": Vertex{17, 20},
	"GoLang": Vertex{20, 20},
	"Rust":   Vertex{13, 20},
}

func compute(fun func(int, int) int, fun2 func(int) int) int {
	return fun(1, 2) + fun2(1)
}

func helper(x int) int { return x }

func main() {

	//....struct literals....//
	// v1 := Vertex{1, 2}
	// v2 := Vertex{X: 3}
	// v3 := Vertex{}
	// p := &Vertex{1, 2}
	// p.X = 10
	// fmt.Println(v1, v2, v3, *p)

	//....Arrays....//
	// var a [2]int
	// a[0] = 1
	// a[1] = 10
	// fmt.Println(a)
	// str := [4]string{"I", "am", "from", "Knackroot"}
	// fmt.Println(str)

	//....Slicing....//
	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// var s []int = primes[:3]
	// arr := [3]int{1, 2, 3}
	// fmt.Println(arr, s)

	//....Slicing reference....//
	// names := [4]string{
	// 	"Alex", "Bob", "Jarvis", "Alice",
	// }
	// n := names[:2]
	// n[0] = "James"
	// fmt.Println(names)

	//....slice literals....//
	// arr := []Vertex{
	// 	{2, 2},
	// 	{5, 2},
	// 	{4, 1},
	// }
	// fmt.Println(arr)
	// newA := arr[:2]
	// fmt.Println(newA)

	//....slice length & capacity....//
	// s := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println(s)
	// s = s[:0]
	// fmt.Println(s)
	// s = s[:4]
	// fmt.Println(s)
	// s = s[3:]
	// fmt.Println(s)
	// s = s[:3]
	// fmt.Printf(s)

	//....nil slices....//
	// var s []int
	// fmt.Println(s)
	// if s == nil {
	// 	fmt.Println("NIL")
	// }

	//....slicing with make function....//
	// a := make([]int, 5)
	// b := make([]int, 0, 5)
	// c := b[:2]
	// d := c[2:5]
	// fmt.Println(a, b, c, d)

	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }
	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"
	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }

	//....Apending Slices....//
	// var s []int
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 0)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 1, 2)
	// fmt.Println(s, len(s), cap(s))

	//....Range....//
	// for i, v := range pow {
	// 	fmt.Println(i, v)
	// }
	// for i := range pow {
	// 	fmt.Println(pow[i])
	// }

	//....Range Continued....//
	// pow := make([]int, 10)
	// for i := range pow {
	// 	pow[i] = 1 << i
	// }
	// fmt.Println(pow)
	// for _, j := range pow {
	// 	fmt.Print(j, " ")
	// }
	// fmt.Println()

	//....Maps....//
	// m := map[string]int{}
	// m["k1"] = 10
	// m["k2"] = 100
	// fmt.Println(m)

	//....Maps....//
	// m := map[string]Vertex{}
	// m["Knackroot"] = Vertex{
	// 	10, 10,
	// }
	// m["Company"] = Vertex{
	// 	20, 100,
	// }
	// fmt.Println(m)
	// fmt.Println(m["Knackroot"])

	//....Maps Literals....//
	// fmt.Println(mp)
	// mp["Google"] = Vertex{10, 200}
	// fmt.Println(mp)

	//....Mutating Maps....//
	// fmt.Println(mp)
	// mp["Java"] = Vertex{2, 10}
	// fmt.Println(mp)
	// delete(mp, "Java")
	// fmt.Println(mp)
	// ok := mp["C++"]
	// if ok.X == 0 && ok.Y == 0 {
	// 	fmt.Println("Not Found")
	// } else {
	// 	fmt.Println(ok)
	// }
	// val, isPresent := mp["Ruby"]
	// fmt.Println(val, isPresent)

	//....Function Values....//
	// add := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Println(add(21, 3))
	// fmt.Println(compute(add, helper))

	//....Function Closures....//
	fun := closure()
	fun("dfsf")
}

func closure() func(string) {
	return func(s string) {
		fmt.Println(s)
	}

	// return f("dsf")
}

func f(s string) {
	fmt.Println(s)
	return
}
