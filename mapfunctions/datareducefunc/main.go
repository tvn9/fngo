package main

import (
	"fmt"
	"fpgo/mapfunctions/datareducefunc/pred"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}

	sum := pred.Sum(ints)
	fmt.Printf("Sum of %v = %d\n", ints, sum)
}
