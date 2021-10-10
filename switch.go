package main

import "fmt"

func main() {
	fmt.Println("---")
	gender := "female"

	switch gender {
	case "male":
		fmt.Println("have pp")
	case "female":
		fmt.Println("no pp")
	default:
		fmt.Println("what are you?")
	}

	// Short statement
	switch length := len(gender); length < 6 {
	case true:
		fmt.Println("have pp")
	case false:
		fmt.Println("no pp")
	}

	switch length := len(gender); length {
	case 4:
		fmt.Println("have pp")
	case 6:
		fmt.Println("no pp")
	}

	// No condition
	length := len(gender)
	switch {
	case length < 6:
		fmt.Println("have pp")
	case length >= 6:
		fmt.Println("no pp")
	}
}
