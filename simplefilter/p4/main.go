package main

import (
	"fmt"
	"fpgo/simplefilter/p4/pkg"
)

type (
	Name   string
	Breed  int
	Gender int
)

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

const (
	Male Gender = iota
	Female
)

func main() {
	dogs := []Dog{
		{"Bucky", Havanese, Male},
		{"Tipsy", Poodle, Female},
	}
	result := pkg.Filter(dogs, func(d Dog) bool {
		return d.Breed == Havanese
	})
	fmt.Printf("Havanese: %v\n", result)

	input := []int{1, 1, 2, 3, 5, 8, 13}
	fmt.Printf("Random numbers from 1 - 30: %v\n", input)

	oddNums := pkg.DropWhile(input, func(i int) bool {
		return i%2 != 0
	})
	fmt.Printf("Odd numbers: %v\n", oddNums)

	oddNum := pkg.TakeWhile(input, func(i int) bool {
		return i%2 != 0
	})
	fmt.Printf("Odd numbers: %v\n", oddNum)
}
