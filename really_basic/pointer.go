package main

import "fmt"

func main() {
	fmt.Println("---")

	addr1 := Address{"Bantul", "Yogyakarta"}
	addr2 := &addr1 // The F*-ing pointer
	addr3 := &addr1

	addr2.City = "Sleman"

	fmt.Println("addr1", addr1)
	fmt.Println("addr2", *addr2) // Accessing value of the F*-ing pointer
	fmt.Println("addr3", *addr3)
	fmt.Println("---")

	addr2.City = "Jogja"

	fmt.Println("addr1", addr1)
	fmt.Println("addr2", *addr2)
	fmt.Println("addr3", *addr3)
	fmt.Println("---")

	addr2 = &Address{"Bandung", "Jabar"} // Doesn't change addr1, only addr2

	fmt.Println("addr1", addr1)
	fmt.Println("addr2", *addr2)
	fmt.Println("addr3", *addr3)
	fmt.Println("---")

	*addr3 = Address{"Solo", "Jateng"} // Change addr1 & addr3 (averyone that point to the same addr as addr3)

	fmt.Println("addr1", addr1)
	fmt.Println("addr2", *addr2)
	fmt.Println("addr3", *addr3)
	fmt.Println("---")

	addr4 := new(Address) // Empty pointer
	fmt.Println(*addr4)
	// addr4.City

}

/**
Golang default = pass by value
*/

type Address struct {
	City, Province string
}
