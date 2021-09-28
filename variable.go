package main

import (
	"fmt"
)

func main() {
	// Cara 1
	var name string

	name = "falih"
	fmt.Println(name)

	name = "falih fauzan"
	fmt.Println(name)

	// Cara 2
	var age = 23
	fmt.Println(age)

	// Cara 3
	gender := "male"
	fmt.Println(gender)

	// Multiple var
	var (
		firsName = "falih"
		lastName = "fauzan"
	)

	fmt.Println(firsName, lastName)
}
