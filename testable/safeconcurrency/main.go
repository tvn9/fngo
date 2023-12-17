// Example of basic safe coccurency
package main

import (
	"fmt"
	"sync"
)

var (
	// interger of type slice of int
	integers = []int{}
)

// addToSlice take a int, and a *sync.WaitGroup, then assign incoming i to slice of intergers
func addToSlice(i int, wg *sync.WaitGroup) {
	integers = append(integers, i)
	wg.Done()
}

// main is the start of go program
func main() {
	// initialize a wg variable
	wg := sync.WaitGroup{}

	numbersToAdd := 10

	wg.Add(numbersToAdd)
	for i := 0; i < numbersToAdd; i++ {
		go addToSlice(i, &wg)
	}
	wg.Wait()
	fmt.Println(integers)
}
