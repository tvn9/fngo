// Type aliases for functions and passing function to functions
package main

import (
	"fmt"
)

// defind a veriable with func(int) bool {} as function signature.
type predicate func(int) bool

// This is where program start => main()
func main() {
	// define a variable that hold the list of numbers
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	largeNums := filter(nums, largerThan8) // passing funtion to to functions

	fmt.Printf("%v\n", largeNums)
}

// Passing function to function
func filter(ns []int, lt8 predicate) []int {
	out := []int{}

	for _, n := range ns {
		if lt8(n) {
			out = append(out, n)
		}
	}
	return out
}

// Regular function
func largerThan8(i int) bool {
	return i > 8
}
