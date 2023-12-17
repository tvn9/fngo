// Example of writing pure function
package main

import "fmt"

/*
var (
	// name is package level variable
	name = "Thanh"
)

// sayHello take reads data from package level variable "name"
// then return a string from "name" as an output
func sayHello() string {
	return fmt.Sprintf("hello %s", name)
}

func main() {
	fmt.Println(sayHello())
}
*/

// Now let rewrite above code in functional style

// sayHello take name variable of type string
// then return a new combined string that copy data fron "name"
func sayHello(name string) string {
	return fmt.Sprintln("hello ", name)
}

func main() {
	fmt.Println(sayHello("Thanh!"))
}
