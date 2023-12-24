package main

import (
	"fmt"
	"fpgo/mapfunctions/dogmapdemo/predicates"
)

func main() {
	dogs := []predicates.Dog{
		{"Backy", predicates.Havanese, predicates.Male},
		{"Tipsy", predicates.Poodle, predicates.Female},
	}
	result := predicates.Map(dogs, func(d predicates.Dog) predicates.Dog {
		if d.Gender == predicates.Male {
			d.Name = "Mr. " + d.Name
		} else {
			d.Name = "Mrs. " + d.Name
		}
		return d
	})
	fmt.Printf("%v\n", result)
}
