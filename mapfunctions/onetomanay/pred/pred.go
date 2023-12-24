package pred

import "fmt"

// Pred is predicate function type
type Pred[A any] func(A) A

// FlatMap
func FlatMap[A any](input []A, m func(A) []A) []A {
	flatout := []A{}

	for _, in := range input {
		ms := m(in)
		flatout = append(flatout, ms...)
	}
	return flatout
}
