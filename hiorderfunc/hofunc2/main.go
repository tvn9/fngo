// Example of higher-order function
// Although we don’t take a function as input, we are returning a function as output.
// This is a valid characteristic of higher-order functions.
package main

import "fmt"

func outerfunc(s string) func() {
	fmt.Println("Outer func: ", s)
	return func() {
		fmt.Println("Inter function")
	}
}

func main() {
	testfunc := outerfunc("Higher-order function")
	testfunc()
}
