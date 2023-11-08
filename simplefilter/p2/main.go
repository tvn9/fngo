// Filter function with Predicate function
package main

import (
	"fmt"
	"math/rand"
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
	// Create 20 numbers slice with random value from 0 to 30
	input := createRandIntSlice(20, 30)
	fmt.Printf("Random number range from 0 to 30: %v\n", input)

	larger20 := Filter(input, func(i int) bool { return i > 20 })
	smaller20 := Filter(input, func(i int) bool { return i < 20 })

	fmt.Printf("Larger > 20 => %v\nSmaller > 20 => %v\n", larger20, smaller20)
}

// //////////////////////////////////////////////////////////////////////////////////////
// These two function below help create the numbers slice and random numbers slice.

/*
// createIntSlice generates a list of 1 to n in sequential order
func createIntSlice(n int) []int {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return nums
}
*/

// createRandIntSlice generates a slice of n length with random value from 1 to n in random order
func createRandIntSlice(n int, rang int) []int {
	if n < 3 || rang < 3 {
		panic("error: please enter value greater than 2")
	}

	nums := []int{}
	for i := 1; i <= n; i++ {
		ran := rand.Intn(rang)
		nums = append(nums, ran)
	}
	return nums
}
