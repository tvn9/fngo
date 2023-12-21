package pred

// predicate function type that only accept type int and return bool
type Predicate func(int) bool

// predicate function that will accept any type of data and return bool
type PredicateA[A any] func(A) bool

// Filter take a slice of int and returns a new int slice with number > 10,
// this is not a flexible function where it is hard code to filter only 10.
func Filter10(numbers []int) []int {
	nums := []int{}

	for _, n := range numbers {
		if n > 10 {
			nums = append(nums, n)
		}
	}
	return nums
}

// FilterLargerThan
func FilterLargerThan(numbers []int, threshold int) []int {
	nums := []int{}

	for _, n := range numbers {
		if n > threshold {
			nums = append(nums, n)
		}
	}
	return nums
}

// FilterSmallerThan
func FilterSmallerThan(numbers []int, threshold int) []int {
	nums := []int{}
	for _, n := range numbers {
		if n < threshold {
			nums = append(nums, n)
		}
	}
	return nums
}

// Filter with PredicateA
func Filter[A any](input []A, p PredicateA[A]) []A {
	output := []A{}
	for _, o := range input {
		if p(o) {
			output = append(output, o)
		}
	}
	return output
}

// Looking for all matches
// AllMatches
func AllMatches[A any](input []A, p PredicateA[A]) bool {
	for _, o := range input {
		if !p(o) {
			return false
		}
	}
	return true
}
