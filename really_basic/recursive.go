package main

import "fmt"

func main() {
	fmt.Println("---")
	fmt.Println(fibo(10))
	fmt.Println(factorial(4))
}

// n-th fibonacci number
func fibo(n int) int {
	if (n == 1) || (n == 2) {
		return 1 // Stopper
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

// n factorial
func factorial(n int) int {
	if n == 1 {
		return 1 // Stopper
	} else {
		return factorial(n-1) * n
	}
}
