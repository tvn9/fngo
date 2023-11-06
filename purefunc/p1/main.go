// Demonstrating pure versus impure function calls
package main

import (
	"fmt"
	"math/rand"
)

// example of a pure function, the return value will
// always be the same with the same input
func add(a, b int) int {
	return a + b
}

// impure function, the return value will be random
// from 1 to 6
func rollDice() int {
	return rand.Intn(6) + 1
}

func main() {

	fmt.Printf("%d + %d = %d\n", 20, 5, add(20, 5))

	for i := 0; i < 5; i++ {
		fmt.Printf("dice roll: %v\n", rollDice())
	}
}
