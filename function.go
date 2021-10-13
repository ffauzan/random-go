package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("---")
	helloWorld(2, "fal")

	fmt.Println(factorSeven(48))

	fmt.Println("---")
	name, length := personDetail("falih")
	fmt.Println(name, length)

	fmt.Println("---")
	_, length2 := personDetail("falih")
	fmt.Println(length2)

	fmt.Println("---")
	user, domain := getEmailDetail("ffauzan@mail.ugm.ac.id")
	fmt.Println("user:", user)
	fmt.Println("domain:", domain)
}

func helloWorld(n int, word string) {
	for i := 0; i < n; i++ {
		fmt.Println("Hello world", word)
	}

}

// With return value
func factorSeven(n int) bool {
	isFactor := false
	if n%7 == 0 {
		isFactor = true
	}
	return isFactor
}

// Return multiple
func personDetail(name string) (string, int) {
	return name, len(name)
}

// Named return value
func getEmailDetail(email string) (user string, domain string) {
	idLimiter := strings.Index(email, "@")
	user = email[:idLimiter]
	domain = email[idLimiter+1:]
	return
}
