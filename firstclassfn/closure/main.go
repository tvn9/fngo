// Capturing variable context in function (closures)
package main

import "fmt"

func main() {
	greetingFunc := createGreeting()
	response := greetingFunc("Thanh")
	fmt.Println(response)

	greetingFunc = createGreeting1("Good morning")
	response = greetingFunc("Thanh")
	fmt.Println(response)
}

func createGreeting() func(string) string {
	s := "Hello "
	return func(name string) string {
		return s + name
	}
}

func createGreeting1(s string) func(string) string {
	return func(name string) string {
		return s + name
	}
}
