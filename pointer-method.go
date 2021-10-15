package main

import "fmt"

func main() {
	fmt.Println("---")
	car1 := Car{"Honda", "Civic"}
	fmt.Println("initial state", car1)

	car1.RemoveBrand1()
	fmt.Println("called RemoveBrand1", car1)

	car1.RemoveBrand2()
	fmt.Println("called RemoveBrand2", car1)
}

type Car struct {
	Brand, Type string
}

// Pass by value
func (car Car) RemoveBrand1() {
	car.Brand = "unknown"
}

// Pass by refference
func (car *Car) RemoveBrand2() {
	car.Brand = "unknown"
}
