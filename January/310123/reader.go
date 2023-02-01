package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("Hello, Reader !")
	// fmt.Println(r)

	b := make([]byte, 8)
	// fmt.Printf("%T", b)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		if err == io.EOF {
			break
		}
	}
}
