package pred

// Predicate function with generic
type Predicate[A any] func(A) bool

// TakeWhile
func TakeWhile[A any](input []A, p Predicate[A]) []A {
	output := []A{}
	for _, o := range input {
		if p(o) {
			output = append(output, o)
		} else {
			return output
		}
	}
	return output
}
