package pred

// Predicate function with generic
type Predicate[A any] func(A) bool

// DropWhile
func DropWhile[A any](input []A, p Predicate[A]) []A {
	output := []A{}
	drop := true
	for _, o := range input {
		if !p(o) {
			drop = false
		}
		if !drop {
			output = append(output, o)
		}
	}
	return output
}
