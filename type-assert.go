package main

import "fmt"

func main() {
	fmt.Println("---")
	res := random()

	resString := res.(string)
	fmt.Println(resString)

	// This will lead to panic becaouse wrong type
	// resInt := res.(int)
	// fmt.Println(resInt)

	// The safe way
	switch value := res.(type) {
	case int:
		fmt.Println(value, "int")
	case string:
		fmt.Println(value, "string")
	default:
		fmt.Println(value, "unknown")
	}
}

func random() interface{} {
	return "OK"
}
