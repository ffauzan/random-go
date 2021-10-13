package main

import "fmt"

func main() {
	fmt.Println("---")
	person1 := newMap("aaa")

	if person1 == nil {
		fmt.Println("Kosong")
	} else {
		fmt.Println(person1)
	}
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
