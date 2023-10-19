// Creating a simple calculator to demonstrade the idea of dispatching function base on certain input value
package main

import "fmt"

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
	return a / b
}

func calculate(a int, operator string, b int) int {

	switch operator {
	case "+":
		return add(a, b)
	case "-":
		return sub(a, b)
	case "*":
		return mult(a, b)
	case "/":
		return div(a, b)
	default:
		panic("operation not supported")
	}
}

type calculateFunc func(int, int) int

var (
	operations = map[string]calculateFunc{
		"+": add,
		"-": sub,
		"*": mult,
		"/": div,
	}
)

func calculateWithMap(a int, operator string, b int) int {
	if operation, ok := operations[operator]; ok {
		return operation(a, b)
	}
	panic("operation not supported.")
}

func main() {
	a := 19
	b := 5

	operator := "/"

	result := calculate(a, operator, b)

	fmt.Printf("%d %s %d = %v\n", a, operator, b, result)

	result1 := calculateWithMap(a, operator, b)

	fmt.Printf("%d %s %d = %v\n", a, operator, b, result1)

}
