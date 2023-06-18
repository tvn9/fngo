package main

import (
	"fmt"
)

type predicate func(int) bool

func main() {
	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	larger := filter(is, largerThan5)

	fmt.Printf("%v\n", larger)
}

func filter(is []int, testNum predicate) []int {
	out := []int{}

	for _, i := range is {
		if testNum(i) {
			out = append(out, i)
		}
	}
	return out
}

func largerThan5(i int) bool {
	return i > 5
}
