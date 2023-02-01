package main

import (
	"fmt"
	"time"
)

func greet(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	go greet("Hey")
	greet("There!")
}
