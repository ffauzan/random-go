package main

import "fmt"

func main() {
	fmt.Println("---")
	helloWorld(2, "fal")

	fmt.Println(factorSeven(48))
}

func helloWorld(n int, word string) {
	for i := 0; i < n; i++ {
		fmt.Println("Hello world", word)
	}

}

// With return value
func factorSeven(n int) bool {
	isFactor := false
	if n%7 == 0 {
		isFactor = true
	}
	return isFactor
}
