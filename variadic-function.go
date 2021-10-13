package main

import "fmt"

func main() {
	fmt.Println("---")
	total := getSum(100, 1, 2, 3, 4, 5)
	fmt.Println(total)

	input := []int{1, 2, 3, 4, 5}
	total2 := getSum(10, input...)
	fmt.Println(total2)
}

func getSum(ratio int, value ...int) int {
	total := 0
	for idx, number := range value {
		total += number
		fmt.Println(idx)
	}

	return total * ratio
}
