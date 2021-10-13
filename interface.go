package main

import "fmt"

func main() {
	fmt.Println("---")

	joko := Person{
		Name: "Joko Widodo",
	}

	munci := Cat{
		Name:   "Munci Uci",
		Gender: "female",
	}

	/**
	Every struct that implement getName() is part of HasName
	Because Person & Cat implementing getName, we can use greet2()
	*/
	greet2(joko)
	greet2(munci)

}

// The interface with getName func
type HasName interface {
	getName() string
}

// The function with HasName interface as an input
func greet2(hasName HasName) {
	fmt.Println("hello", hasName.getName())
}

// The struct
type Person struct {
	Name string
}

type Cat struct {
	Name, Gender string
}

// Person struct method implementing getName()
func (person Person) getName() string {
	return person.Name
}

// Cat struct method implementing getName()
func (cat Cat) getName() string {
	if cat.Gender == "male" {
		return "Mr." + cat.Name
	} else {
		return "Ms. " + cat.Name
	}
}
