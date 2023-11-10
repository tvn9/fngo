package predicates

import "fmt"

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

type (
	Name   string
	Breed  int
	Gender int
)

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

type MapFunc[A any] func(A) A

func Map[A any](input []A, m MapFunc[A]) []A {
	output := make([]A, len(input))
	for i, element := range input {
		output[i] = m(element)
	}
	return output
}

func DogMapDemo() {
	dogs := []Dog{
		{"Backy", Havanese, Male},
		{"Tipsy", Poodle, Female},
	}
	result := Map(dogs, func(d Dog) Dog {
		if d.Gender == Male {
			d.Name = "Mr. " + d.Name
		} else {
			d.Name = "Mrs. " + d.Name
		}
		return d
	})
	fmt.Printf("%v\n", result)
}
