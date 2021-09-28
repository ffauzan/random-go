package main

import (
	"fmt"
)

func main() {
	var a = 80
	var b = 10
	var c = a + b

	fmt.Println(c)

	// Augmented assignment
	c += 10
	fmt.Println(c)

	// Unary
	c++
	fmt.Println(c)

}
