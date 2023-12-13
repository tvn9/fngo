// Example of Partial application
package main

import "fmt"

func createGreeting(g string) func(s string) string {
	return func(name string) string {
		return g + name
	}
}

func main() {
	firstGreeting := createGreeting("Well, hello there ")
	secondGreeting := createGreeting("Hola ")
	fmt.Println(firstGreeting("Thanh"))
	fmt.Println(firstGreeting("Xuan"))
	fmt.Println(secondGreeting("Kim"))
}
