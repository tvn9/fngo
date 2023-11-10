package main

import (
	"fmt"

	"github.com/tvn9/fpgo/reduce/utils"
)

func main() {
	// Sum
	ints := []int{1, 2, 3, 4}
	sum := utils.Sum(ints)
	fmt.Printf("%v\n", sum)

	// Product
	product := utils.Product(ints)
	fmt.Printf("%v\n", product)

	//
	words := []string{"hello", "world", "universe"}
	result := utils.ReduceWithStart(words, "first", func(s1, s2 string) string {
		return s1 + ", " + s2
	})
	fmt.Printf("%v\n", result)
}
