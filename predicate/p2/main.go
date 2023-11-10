package main

import (
	"fmt"

	"github.com/tvn9/fpgo/predicate/p2/predicate"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	result := predicate.FlatMap(ints, func(n int) []int {
		out := []int{}
		for i := 0; i < n; i++ {
			out = append(out, i)
		}
		return out
	})
	fmt.Printf("%v\n", result)
}
