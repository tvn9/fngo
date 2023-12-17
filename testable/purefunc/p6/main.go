// Improved confidence in fucntion names and signatures
package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		i = add1(i)
		fmt.Printf("%v, ", i)
	}
}

func add1(input int) int {
	if input == 10 {
		panic("can not increment any more.")
	}
	return input + 1
}
