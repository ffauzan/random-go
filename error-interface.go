package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("---")
	res, err := devide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func devide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot devide by zero")
	} else {
		return a / b, nil
	}
}
