package main

import "fmt"

func main() {
	fmt.Println("---")

	x := 1
	y := 2

	fmt.Println(x, y)

	switchVal(&x, &y)

	fmt.Println(x, y)
}

func switchVal(a *int, b *int) {
	*a = 10
	*b = 20
}
