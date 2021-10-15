package main

import (
	"fmt"
)

// Type declaration
type checker func(string) string

func main() {
	fmt.Println("---")

	sayHello("fal", genderCheck, "male")
	sayHello2("fal", genderCheck, "male")

}

func genderCheck(gender string) string {
	if gender == "male" {
		return "mas"
	} else if gender == "female" {
		return "mbak"
	} else {
		return "lur"
	}
}

func sayHello(name string, filter func(string) string, gender string) {
	fmt.Println("hello", filter(gender), name)
}

// With type declaration
func sayHello2(name string, filter checker, gender string) {
	fmt.Println("hello", filter(gender), name)
}
