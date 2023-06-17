package main

import "fmt"

type predicate func(i int) bool

func main() {
	// Anonymous function
	larger := filter([]int{1, 2, 3, 4, 5, 6, 7}, func(i int) bool { return i > 2 })

	fmt.Println(larger)

}

func filter(ints []int, p predicate) []int {
	out := []int{}
	for _, i := range ints {
		if p(i) {
			out = append(out, i)
		}
	}
	return out
}
