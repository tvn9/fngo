// Introduction to higher-order functions
package main

import "fmt"

func Hello() string {
	return "Hello,"
}

func World() string {
	return Hello() + " world!"
}

func main() {
	fmt.Println(World())
}
