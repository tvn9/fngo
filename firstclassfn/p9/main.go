// Store function inside a map
package main

import "fmt"

type predicate func(i int) bool

var (
	largerThanTwo     = createLargerThanPredicate(2)
	largerThanFive    = createLargerThanPredicate(5)
	largerThanHundred = createLargerThanPredicate(100)
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 9, 100, 102}

	dispatcher := map[string]predicate{
		"2":   largerThanTwo,
		"5":   largerThanFive,
		"100": largerThanHundred,
	}

	fmt.Printf("%v\n", filter(ints, dispatcher["100"]))
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
