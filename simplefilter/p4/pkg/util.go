package pkg

import "math/rand"

// //////////////////////////////////////////////////////////////////////////////////////
// These two function below help create the numbers slice and random numbers slice.

// createIntSlice generates a list of 1 to n in sequential order
func CreateIntSlice(from, too int) []int {
	nums := []int{}
	for i := from; i <= too; i++ {
		nums = append(nums, i)
	}
	return nums
}

// createRandIntSlice generates a slice of n length with random value from 1 to n in random order
func CreateRandIntSlice(n int, rang int) []int {
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
