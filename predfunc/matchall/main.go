package main

import (
	"fmt"
	"fpgo/predfunc/matchall/pred"
)

func main() {
	input := []int{2, 4, 5, 8, 9, 15, 55, 55, 55, 3, 35, 21, 68, 99}
	all5 := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}

	// Hardcode filter to 10
	nums := pred.Filter10(input)
	fmt.Printf("Number > 10: %d\n", nums)

	// Filter larger than threshold
	larger20 := pred.FilterLargerThan(input, 20)
	fmt.Printf("Number > 20: %d\n", larger20)

	// Filter smaller than threshold
	smaller10 := pred.FilterSmallerThan(input, 10)
	fmt.Printf("Number < 10: %d\n", smaller10)

	// Using the Filter with Predicate and Generic features
	larger15 := pred.Filter(input, func(i int) bool { return i > 15 })
	fmt.Printf("Number > 15 %d\n", larger15)

	// Let try largerthan and smallerthan
	larger5smaller20 := pred.Filter(input, func(i int) bool {
		return i > 5 && i < 20
	})
	fmt.Printf("Number < 5 and > 20: %d\n", larger5smaller20)

	// Find a match
	filtered := pred.Filter(input, func(i int) bool { return i == 55 })
	fmt.Printf("Matching 55: %d\n", filtered)
	contains55 := len(filtered) > 0
	fmt.Printf("Number of mactches %d, %v\n", len(filtered), contains55)

	// Looking for all matches
	allmatches := pred.AllMatches(all5, func(i int) bool { return i == 5 })
	fmt.Printf("Input: %d\n", all5)
	fmt.Printf("All matches %v\n", allmatches)

}
