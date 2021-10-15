package main

import "fmt"

func main() {
	fmt.Println("---")
	helloPeople := hello
	fmt.Print(helloPeople("falfal"))
}

func hello(name string) string {
	return "hello " + name
}
