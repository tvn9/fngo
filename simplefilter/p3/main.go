// Filter function with Predicate function
package main

import (
	"fmt"

	"github.com/tvn9/fpgo/simplefilter/p3/pkg"
)

type Predicate[A any] func(A) bool

// Filter (with generic, and first class function)
func Filter[A any](input []A, pred Predicate[A]) []A {
	output := []A{}
	for _, element := range input {
		if pred(element) {
			output = append(output, element)
		}
	}
	return output
}

func main() {
	// Credit a sequential order numbers slice rang from and to n
	test := pkg.CreateIntSlice(10, 30)
	fmt.Println(test)

	// Create 20 numbers slice with random value from 0 to 30
	input := pkg.CreateRandIntSlice(20, 30)
	fmt.Printf("Random number range from 0 to 30: %v\n", input)

	// Filter numbers larger than 20 or smaller than 20
	larger20 := Filter(input, func(i int) bool { return i > 20 })
	smaller20 := Filter(input, func(i int) bool { return i < 20 })

	fmt.Printf("Larger > 20 => %v\nSmaller > 20 => %v\n", larger20, smaller20)

	// Filter numbers > 10 and < 20
	larger10smaller20 := Filter(input, func(i int) bool { return i > 10 && i < 20 })
	fmt.Printf("Numbers > 10 and < 20 => %v\n", larger10smaller20)

	// Filter even numbers
	evenNumber := Filter(input, func(i int) bool { return i%2 == 0 })
	fmt.Printf("Even numbers => %v\n", evenNumber)
}
