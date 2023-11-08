// Implementing a filter function
package main

// Filter with fix set threshold
func Filter(numbers []int) []int {
	out := []int{}
	for _, num := range numbers {
		if num > 10 {
			out = append(out, num)
		}
	}
	return out
}

// Filter with allow user to set threshold
func FilterTheshold(numbers []int, threshold int) []int {
	out := []int{}
	for _, num := range numbers {
		if num > threshold {
			out = append(out, num)
		}
	}
	return out
}
