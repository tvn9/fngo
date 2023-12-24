package main

import (
	"fmt"
	"fpgo/mapfunctions/mapfunc/mMap"
)

func main() {
	ints := []int{1, 1, 2, 3, 5, 8, 13}
	time2 := mMap.Map(ints, func(i int) int {
		return i * 2
	})
	fmt.Printf("%v\n", time2)
}
