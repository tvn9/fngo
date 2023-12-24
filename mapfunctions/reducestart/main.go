package main

import (
	"fmt"
	"fpgo/mapfunctions/reducestart/pred"
)

func main() {
	words := []string{"hello", "world", "universe"}
	result := pred.ReduceWithStart(words, "first", func(s1, s2 string) string {
		return s1 + ", " + s2
	})
	fmt.Printf("%v\n", result)
}
