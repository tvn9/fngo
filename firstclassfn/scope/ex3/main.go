// Scoping example 3
package main

import "fmt"

func main() {
	s := "hello" // "hello" string is assigned to s
	if true {
		s = "world" // "world" string is assigned to s
		fmt.Println(s)
	}
	fmt.Println(s)
}

// The out put will be "world world"
