// package main

// import "fmt"

// func f(arr []int, c chan int) {

// 	sum := 0

// 	for _, v := range arr {
// 		sum += v
// 	}

// 	c <- sum
// }

// func main() {

// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(arr)

// 	c := make(chan int)
// 	go f(arr[:5], c)
// 	go f(arr[5:], c)

// 	x, y := <-c, <-c
// 	fmt.Println(x, y)

// 	ch := make(chan int, 3)
// 	ch <- 2
// 	ch <- 3
// 	ch <- 4

// 	a := <-ch
// 	fmt.Println(a)
// 	ch <- 5
// }

package main

import (
	"fmt"
	"time"
)

func myfunc(ch chan int) {

	time.Sleep(5 * time.Second)
	fmt.Println(<-ch)
}

func main() {

	fmt.Println("start Main")

	ch := make(chan int)

	go myfunc(ch)

	fmt.Println(1, "End Main")
	ch <- 23
	fmt.Println(2, "End Main")

}
