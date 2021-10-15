package main

import (
	"fmt"
)

func main() {
	var n32 int32 = 300
	var n8 int8 = int8(n32)

	var name string = "falFal"

	fmt.Println(n32)
	fmt.Println(n8)

	//STR array of char
	fmt.Println(name[3])
	fmt.Println(string(name[3]))
}
