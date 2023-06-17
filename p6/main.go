// Returning functions from functions
package main

import "fmt"

type predicate func(i int) bool

func createLargerThanPredicate(threadhold int) predicate {
	return func(i int) bool {
		return i > threadhold
	}
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 101, 120, 234}

	largerThanTwo := createLargerThanPredicate(2)
	largerThanFive := createLargerThanPredicate(5)
	largerThanHundred := createLargerThanPredicate(100)

	twoPlus := filter(ints, largerThanTwo)
	fivePlus := filter(ints, largerThanFive)
	hundredPlus := filter(ints, largerThanHundred)

	fmt.Println(twoPlus)
	fmt.Println(fivePlus)
	fmt.Println(hundredPlus)
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
