package main

import "fmt"

func main() {
	fmt.Println("---")

	// Break = break the f-ing loop, even the post statement is not executed (i++)
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// Continue = ignore all code below and go to next iteration
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i, "even")
	}

}
