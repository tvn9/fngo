// Predicate is a statement that can be evaluated as either true or false.
package main

import (
	"fmt"
	"fpgo/predfunc/filter/preda"
)

func main() {
	input := []int{1, 1, 3, 5, 8, 13, 21, 34, 55}
	larger0smaller20 := preda.FilterGeneric(input, func(i int) bool {
		return i > 10 && i < 20
	})
	evenNumbers := preda.FilterGeneric(input, func(i int) bool {
		return i%2 == 0
	})
	fmt.Printf("%v\n%v\n", larger0smaller20, evenNumbers)
}
