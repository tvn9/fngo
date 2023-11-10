package predicate

func FlatMap[A any](input []A, m func(A) []A) []A {
	output := []A{}
	for _, element := range input {
		newElements := m(element)
		output = append(output, newElements...)
	}
	return output
}
