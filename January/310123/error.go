package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

type E struct {
	s string
}

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s",
// 		e.When e.What)
// }

func (e E) Error() string {
	e.s = "Edsfsd"
	return e.s
}

func run() error {
	// return &MyError{
	// 	time.Now(),
	// 	"it didn't work",
	// }

	return E{
		"Not working",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		fmt.Printf("%T", err)
	}

	t := MyError{time.Now(), "dfsdf"}

	fmt.Printf("%T", t.When)
}
