// Writing immutable code for collection data types
package main

import "fmt"

func main() {
	m := map[string]int{}
	addValue(m, "blue", 6)
	addValue(m, "Brown", 1)
	fmt.Printf("%v\n", m)
}

func addValue(m map[string]int, s string, n int) {
	m[s] = n
}
