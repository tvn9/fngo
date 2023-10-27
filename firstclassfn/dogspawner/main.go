// Example: DogSpawner
package main

import "fmt"

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

type (
	Name          string
	Breed         int
	Gender        int
	NameToDogFunc func(Name) Dog
)

// define possible breeds
const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

// define possible genders
const (
	Male Gender = iota
	Female
)

var (
	maleHavaneseSpawner = DogSpawner(Havanese, Male)
	femalePoodleSpawner = DogSpawner(Poodle, Female)
)

/*
func createDogsWithoutPartialApplicaiton() {
	bucky := Dog{
		Name:   "Bucky",
		Breed:  Havanese,
		Gender: Male,
	}
	rocky := Dog{
		Name:   "Rocky",
		Breed:  Havanese,
		Gender: Male,
	}
	tipsy := Dog{
		Name:   "Tipsy",
		Breed:  Poodle,
		Gender: Female,
	}
}
*/

func DogSpawner(b Breed, g Gender) NameToDogFunc {
	return func(n Name) Dog {
		return Dog{
			Breed:  b,
			Gender: g,
			Name:   n,
		}
	}
}

func main() {
	bucky := maleHavaneseSpawner("bucky")
	rocky := maleHavaneseSpawner("rocky")
	tipsy := femalePoodleSpawner("tipsy")
	fmt.Printf("%v\n", bucky)
	fmt.Printf("%v\n", rocky)
	fmt.Printf("%v\n", tipsy)
}
