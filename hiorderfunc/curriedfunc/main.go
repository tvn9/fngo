// Function currying, and how to reduce n-ary function to unary function
package main

import "fmt"

// threeSum return the sum of a, b, and c variables
// this is the regular function
func threeSum(a, b, c int) int {
	return a + b + c
}

// threeSumCurried is the currying function that also return the sum of a, b, and c
func threeSumCurried(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

// main - start the program
func main() {
	fmt.Println(threeSum(20, 30, 10))
	fmt.Println(threeSumCurried(20)(30)(10))
}
