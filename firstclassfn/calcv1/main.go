// Refacturing the previous calculator program to functional style
package main

import (
	"fmt"
)

// calculateFunc setup a variable with type func(int, int) int
type calculateFunc func(int, int) int

// operators maps the calculator operation string to the corressponding function
var (
	operators = map[string]calculateFunc{
		"+": add,
		"-": sub,
		"*": mult,
		"/": div,
	}
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		panic("devide by zero")
	}
	return (a / b)
}

// calculateWithMap perform calculation using map dispatcher pattern
func calculateWithMap(a, b int, op string) int {
	if operator, ok := operators[op]; ok {
		return operator(a, b)
	}
	return 0
}

// The good old main
func main() {
	a := 19
	b := 5
	operator := "*"
	result := calculateWithMap(a, b, operator)
	fmt.Printf("%d %s %d = %v\n", a, operator, b, result)
}
