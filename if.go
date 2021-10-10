package main

import "fmt"

func main() {
	// Normal IF-Else
	status := "jomblo"
	if status == "married" {
		fmt.Println("mantab jiwa")
	} else if status == "jomblo" {
		fmt.Println("aduh")
	} else {
		fmt.Println("gjls")
	}

	// If short statement
	if n := len(status); n > 5 {
		fmt.Println("nama bagus")
	}
}
