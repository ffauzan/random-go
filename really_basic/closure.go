package main

import "fmt"

func main() {
	fmt.Println("---")

	name1 := "falfal"
	name2 := "lihlih"

	thisFunction := func() {
		// this will change
		name1 = "not falfal"
	}

	anotherFunction := func() {
		// re declare, doesn't change
		name2 := "not lihlih"
		name2 += "huhu"
	}

	thisFunction()
	anotherFunction()

	fmt.Println(name1)
	fmt.Println(name2)
}
