package main

import (
	"fmt"
)

func main() {
	// Slice = part of array
	// Actually referencing to the array (change in array -> change in slice and vice versa)
	// Property of slice: pointer(pointing at 1st data), length, capacity
	// length <= capacity
	// capacity: quantity from pointer to end of original array
	// arr[low:high] -> create slice start from index low to before index high (index high - 1)
	// ex: arr[4:7] -> p:4, l:3, c:8()

	var (
		// data1  = [10]int{1, 2, 3, 4, 5, 6}
		months = [...]string{
			"jan",
			"feb",
			"mar",
			"apr",
			"mei",
			"jun",
			"jul",
			"aug",
			"sep",
			"oct",
			"nov",
			"des",
		}
		slice1 = months[2:8]
	)

	fmt.Println(slice1)
	fmt.Println(months[6:])
	fmt.Println(months[:6])

	// Changing slice and array
	slice1[0] = "ahahaha"
	fmt.Println(months)

	months[2] = "ahihihi"
	fmt.Println(slice1)

	fmt.Println(cap(months[2:8]))

	// Append
	// Case1: append till over cap
	// it will create new slice, so change in silce2 will not alter months
	slice2 := append(months[:], months[11:]...)
	slice2[0] = "anjay"
	fmt.Println("-------------------append 1--------------------")
	fmt.Println(slice2)
	fmt.Println(months)

	// Case2: append within cap
	// it will NOT create new slice, so change in silce2 will ALTER months
	slice3 := append(months[:1], months[11:]...)
	slice3[0] = "anjay manjay"
	fmt.Println("-------------------append 2--------------------")
	fmt.Println(slice3)
	fmt.Println(months)

	// Create new slice
	newSlice := make([]string, 4, 5)

	newSlice[0] = "zero"
	newSlice[1] = "one"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	// Array vs Slice declaration
	thisArray := [...]int{1, 2, 3, 4}
	thisSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisArray)

	fmt.Println(len(thisSlice))
	fmt.Println(cap(thisSlice))
}
