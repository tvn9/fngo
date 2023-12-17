// Referential Transparancy
package main

import "fmt"

func main() {
	fmt.Printf("%v\n", add(10, add(10, 5)))
	fmt.Printf("%v\n", add(10, 15))
}

func add(a, b int) int {
	return a + b
}
