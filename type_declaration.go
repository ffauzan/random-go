package main

import (
	"fmt"
)

func main() {
	type nama string
	type status bool

	var namaDepan nama = "anjayman"
	var isOpen status = true

	fmt.Println(namaDepan)
	fmt.Println(isOpen)
}
