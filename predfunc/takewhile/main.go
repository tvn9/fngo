package main

import (
	"fmt"
	"fpgo/predfunc/takewhile/pred"
)

func main() {
	input := []int{2, 4, 5, 8, 9, 15, 50, 26, 55, 3, 35, 21, 68, 99}
	// all5 := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}

	evenNums := pred.TakeWhile(input, func(i int) bool { return i%2 == 0 })
	fmt.Printf("Odd numbers: %d\n", evenNums)
}
