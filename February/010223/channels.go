package main

import "fmt"

func myfunc(c chan int) {
	fmt.Println(10 + <-c)
}

func f(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	c := make(chan int)
	go myfunc(c)
	c <- 10
	fmt.Println("End")

	ch := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	go f(arr[:len(arr)/2], ch)
}
