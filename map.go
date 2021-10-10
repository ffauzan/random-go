package main

import (
	"fmt"
)

func main() {
	fmt.Println("anjay")

	country := map[string]string{
		"ID": "Indonesia",
		"IN": "India",
		"US": "Amerika",
	}
	country2 := map[string]string{}

	country["CA"] = "Canada"

	fmt.Println(country["IN"])
	country["IN"] = "Adudu"
	fmt.Println(country["IN"])

	delete(country, "ID")
	fmt.Println(len(country2))

}
