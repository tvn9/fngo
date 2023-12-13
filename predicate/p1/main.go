// Predicate is a statement that can be evaluated as either true or false.
package main

import (
	"fmt"

	"github.com/tvn9/fpgo/predicate/p1/predicates"
)

func main() {
	input := []int{1, 1, 3, 5, 8, 13, 21, 34, 55}
	larger0smaller20 := predicates.FilterGeneric(input, func(i int) bool {
		return i > 10 && i < 20
	})
	evenNumbers := predicates.FilterGeneric(input, func(i int) bool {
		return i%2 == 0
	})
	fmt.Printf("%v\n%v\n", larger0smaller20, evenNumbers)
}













