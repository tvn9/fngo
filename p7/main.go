// Function inside data structures
package main

import "fmt"

var (
	largerThanTwo     = createLargerThanPredicate(2)
	largerThanFive    = createLargerThanPredicate(5)
	largerThanHundred = createLargerThanPredicate(100)
)

type predicate func(i int) bool

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 101, 120, 234}
	predicates := []predicate{largerThanTwo, largerThanFive, largerThanHundred}

	for _, predicate := range predicates {
		fmt.Printf("%v\n", filter(ints, predicate))
	}
}

func createLargerThanPredicate(threshold int) predicate {
	return func(i int) bool {
		return i > threshold
	}
}

func filter(ints []int, p predicate) []int {
	var out []int
	for _, i := range ints {
		if p(i) {
			out = append(out, i)
		}
	}
	return out
}
