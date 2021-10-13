package main

import "fmt"

func main() {
	fmt.Println("---")

	runThing(0)
}

func logging() {
	fmt.Println("Finished running function")
}

func runThing(n int) {
	defer logging() // logging still called even if runThing got error
	n = 10 / n
	fmt.Println(n)
}
