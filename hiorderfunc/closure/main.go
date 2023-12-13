// Example of closure, a closure is any inner function that use
// a variable introduced in the outer function to perform its work.
package main

import "fmt"

func ask() func(string) string {
	s := "What is your name? "
	return func(name string) string {
		return s + name
	}
}

func main() {
	interview := ask()
	name := interview("Thanh Vu Nguyen.")
	fmt.Println(name)
}
