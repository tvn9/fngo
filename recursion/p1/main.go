package main

import "fmt"

func main() {
	result1 := fact(5)
	fmt.Println(result1)
	result2 := factorial(5)
	fmt.Println(result2)
}

func fact(input int) int {
	if input == 0 {
		return 1
	}
	return input * fact(input-1)
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}
