package main

import "fmt"

type Blacklist func(string) bool

func main() {
	fmt.Println("---")

	checker := func(name string) bool {
		if name == "falfal" {
			return false
		} else {
			return true
		}
	}

	register("falfal", checker)
}

func register(name string, isblacklist Blacklist) {
	if isblacklist(name) {
		fmt.Println("User:", name, "is blocked")
	} else {
		fmt.Println("welcome", name)
	}
}
