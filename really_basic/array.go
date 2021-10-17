package main

import (
	"fmt"
)

func main() {
	var data [10]string
	data[0] = "cah 0"

	var data2 = [4]int{
		1,
		2,
		3,
	}

	fmt.Println(data2[0])
	fmt.Println(len(data2))
}
