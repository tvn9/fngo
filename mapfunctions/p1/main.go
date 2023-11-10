// Predicate-based functions
package main

import "fmt"

// Example of a predicate function and generic
type Mapfunc[A any] func(A) A

func Map[A any](input []A, m Mapfunc[A]) []A {
	output := make([]A, len(input))
	for i, element := range input {
		output[i] = m(element)
	}
	return output
}

func main() {
	ints := []int{1, 1, 2, 3, 5, 8, 13}
	result := Map(ints, func(i int) int {
		return i * 2
	})
	fmt.Printf("%v\n", result)
}
