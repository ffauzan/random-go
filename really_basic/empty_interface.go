package main

import "fmt"

func main() {
	fmt.Println("---")
	var value1 interface{} = getAny(1)
	value2 := getAny(0)
	/**
	you can't do this:
	var value int = getAny(1)
	even if the return for that case is int
	*/

	fmt.Println(value1)
	fmt.Println(value2)
}

/**
Example of empty interface
fmt.Println(a ...interface{})
so we can pass anything as arguments to fmt.Println()
*/

type any interface{}

/**
Use Case:
create function that can return any type of data
to catch re return value, you should define the variable as interface{}
*/
func getAny(i int) interface{} {
	switch i {
	case 0:
		return "zerooooooo"
	default:
		return i
	}
}
