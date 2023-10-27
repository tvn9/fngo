// Scoping example 4:
package main

import "fmt"

func main() {
	s := "hello"
	// s := "world" // Error: s variable is already existed.
	// This will resulted in complie error as Go does not
	// allow reuse of variable name in the same block scope

	fmt.Println(s)
}
