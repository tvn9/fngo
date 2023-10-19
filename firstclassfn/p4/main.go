// In-line function
package main

import "fmt"

type predicate func(int) bool

func main() {
	// inline struct
	inlinePersonStruct := struct {
		name string
	}{
		name: "Thanh",
	}

	// ints slice
	ints := []int{1, 2, 3, 5, 6, 7}

	// inline function
	inlineFunction := func(i int) bool {
		return i > 2
	}

	larger := filter(ints, inlineFunction)

	fmt.Println(inlinePersonStruct.name)
	fmt.Println(larger)
}

// passing function to function
func filter(ints []int, p predicate) []int {
	out := []int{}
	for _, i := range ints {
		if p(i) {
			out = append(out, i)
		}
	}
	return out
}
