package main

import "fmt"

func main() {
	fmt.Println("---")

	// Normal way
	var budi User
	budi.Id = 1
	budi.Type = 1
	budi.Name = "Budi Widodo"

	// Literals
	bambang := User{
		Id:   2,
		Type: 1,
		Name: "Bambang Widodo",
	}

	/**
	Short but not good for long run
	If you change position/delete/add field it will f*cked up
	*/
	joko := User{3, 2, "Joko Widodo"}

	fmt.Println(budi)
	fmt.Println(bambang)
	fmt.Println(joko)

	/**
	What really happen is:
	call greet() with joko as an argument
	*/
	joko.greet("morning")

}

// Normal way
type User struct {
	Id, Type int
	Name     string
}

// Struct method (not really, cz struct ahs no behavior)
func (u User) greet(time string) {
	fmt.Println("Hello", u.Name, ", good", time)
}
