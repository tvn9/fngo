package preda

// Predicate function type with go generic
type Predicate[A any] func(A) bool

// Filter returns any number greater than 10
func Filter(numbers []int) []int {
	out := []int{}
	for _, num := range numbers {
		if num > 10 {
			out = append(out, num)
		}
	}
	return out
}

// FilterWithThreshold allow user to set the test value and return any
// value that match the setting condition.
func FilterWithThreshold(numbers []int, threshold int) []int {
	out := []int{}
	for _, num := range numbers {
		if num > threshold {
			out = append(out, num)
		}
	}
	return out
}

// FilterGeneric allows the input of any data type
func FilterGeneric[A any](input []A, pred Predicate[A]) []A {
	output := []A{}
	for _, element := range input {
		if pred(element) {
			output = append(output, element)
		}
	}
	return output
}
