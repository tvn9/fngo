// Scoping example 2
package main

import "fmt"

func main() {
	s := "hello"
	if true {
		s := "world"
		fmt.Println(s)
	}
	fmt.Println(s)
}

// Out put will be: world hello becase s variable
// main and s variable in the if block are too different
// variables, as Go allows using same variable name as
// long as they are in different block scope.
