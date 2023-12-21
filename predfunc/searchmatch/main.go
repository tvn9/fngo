package main

import (
	"fmt"
	"fpgo/predfunc/searchmatch/pred"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 13, 23, 55, 33}
	result := pred.FlatMap(ints, func(n int) []int {
		out := []int{}
		for i := 0; i < n; i++ {
			out = append(out, i)
		}
		return out
	})
	fmt.Printf("%v\n", result)
}
