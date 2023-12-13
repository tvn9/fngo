package main

import (
	"fmt"
)

type predicate func(i int) bool

func LargerThan(filter int) predicate {
	return func(i int) bool { return i > filter }
}

func SmallerThan(filter int) predicate {
	return func(i int) bool { return i < filter }
}

var (
	larger5  = LargerThan(5)
	larget45 = LargerThan(45)
	smaller5 = SmallerThan(5)
)

func main() {
	nums := []int{1, 32, 2, 5, 43, 45, 96, 8, 33}

	predicates := []predicate{larger5, larget45, smaller5}

	for _, pred := range predicates {
		out := Filter(nums, pred)
		fmt.Println(out)
	}
}

func Filter(nums []int, p predicate) []int {
	var out []int
	for _, n := range nums {
		if p(n) {
			out = append(out, n)
		}
	}
	return out
}

/*
func main() {
	nums := []int{1, 32, 2, 5, 43, 45, 96, 8, 33}

	filter := func(n []int, p func(i int) bool) []int {
		var out []int
		for _, n := range nums {
			if p(n) {
				out = append(out, n)
			}
		}
		return out
	}

	out := filter(nums, func(i int) bool { return i%2 == 0 })

	fmt.Println(out)

}
*/

/*
	func main() {
		str1, err := func1()
		if err != nil {
			panic(err)
		}

		str2, err := func2()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", str1)
		fmt.Printf("%v\n", str2)
	}

	func func1() (string, error) {
		return "Thanh", errors.New("error 1")
	}

	func func2() (string, error) {
		return "Nguyen", errors.New("error 2")
	}
*/

/*
type predicate func(i int) bool

func Filter(nums []int, p predicate) []int {
	var out []int
	for _, n := range nums {
		if p(n) {
			out = append(out, n)
		}
	}
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// find even numbers
	evenNum := func(i int) bool {
		if i > 0 {
			if i%2 == 0 {
				return true
			}
		}
		return false
	}

	fmt.Println(Filter(nums, evenNum))
}

*/
