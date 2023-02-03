package main

import "fmt"

func main() {

	ch := make(chan int)
	quit := make(chan int)

	go func() {
		i := 0
		for {
			if i == 10 {
				break
			}
			fmt.Println(<-ch)
			i++
		}
		quit <- 0
	}()

	fibo(ch, quit)
}

func fibo(ch, quit chan int) {

	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
