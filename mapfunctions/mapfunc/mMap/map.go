package mMap

// MapFunc is type definition for a generic function
type MapFunc[A any] func(A) A

func Map[A any](input []A, m MapFunc[A]) []A {
	// create a slice with specific len dinamically assign using the length of the input
	output := make([]A, len(input))

	for i, e := range input {
		output[i] = m(e)
	}
	return output
}
