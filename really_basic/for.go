package main

import "fmt"

func main() {
	fmt.Println("---")

	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// With statement
	// init statement; condition; post statement
	for j := 0; j < 5; j++ {
		fmt.Println(j, "J")
	}

	// For range (work on array, slice, map)
	slice := []string{
		"anjay",
		"wadidaw",
		"which is",
		"adudu",
	}

	for index, value := range slice {
		fmt.Println(index, value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

}
