package main

import (
	"fmt"
)

func main() {
	var (
		a = true
		b = false
		c bool
	)

	c = (a && b) || !b

	fmt.Print(c)

}
