package main

import (
	"fmt"
	"fpgo/predfunc/dropwhile/pred"
)

func main() {
	input := []int{2, 4, 5, 8, 9, 15, 55, 55, 55, 3, 35, 21, 68, 99}
	// all5 := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}

	evenNums := pred.DropWhile(input, func(i int) bool { return i%2 != 0 })
	fmt.Printf("Even numbers: %d\n", evenNums)
}
