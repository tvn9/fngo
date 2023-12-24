package main

import "fmt"

// main starts the program
func main() {
	result1 := fact(5)
	fmt.Println(result1)
	result2 := factorial(5)
	fmt.Println(result2)
}

// fact is a recursive factorial function
func fact(input int) int {
	if input == 0 {
		return 1
	}
	return input * fact(input-1)
}

// factorial in iterative form
func factorial(n int) int {
	num := 1
	for i := 1; i <= n; i++ {
		num *= i
		fmt.Printf("%d * %d\n", i, num)
	}
	fmt.Println(num)
	return num
}
