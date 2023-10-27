package main

import (
	"errors"
	"fmt"
)

func main() {
	str1, err := func1()
	if err != nil {
		panic(err)
	}

	str2, err := func2()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", str1)
	fmt.Printf("%v\n", str2)
}

func func1() (string, error) {
	return "Thanh", errors.New("error 1")
}

func func2() (string, error) {
	return "Nguyen", errors.New("error 2")
}
