package main

import "fmt"

func main() {
	fmt.Println("---")
	runApp2(true)
	fmt.Println("---")
}

func endApp() {
	fmt.Println("Application stopped")
}

// Panic to stop the app
func runApp(err bool) {
	defer endApp()
	if err {
		panic("OMG we got an error")
	}
	fmt.Println("aplication running")
}

// Using recover to keep the program running
func runApp2(err bool) {
	defer endApp2()
	if err {
		panic("OMG we got an error")
	}
	fmt.Println("aplication running")
}

/**
Use recover in defer
*/
func endApp2() {
	err := recover()
	if err != nil {
		fmt.Println("got panic with:", err)
	}

	fmt.Println("Application stopped")
}
