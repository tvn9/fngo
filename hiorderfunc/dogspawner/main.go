// Example of DogSpawner app
package main

import "fmt"

// type alisa for primitive
type (
	Name          string
	Breed         byte
	Gender        byte
	NameToDogFunc func(Name) Dog
)

// define breeds
const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

// define genders
const (
	Male Gender = iota
	Female
)

// Define Dog
type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

/*
func createDogsWithoutPartialApplicaiton() {
	bucky := Dog{
		Name:   "Bucky",
		Breed:  Havanese,
		Gender: Male,
	}
	rocky := Dog{
		Name:   "Rockey",
		Breed:  Havanese,
		Gender: Male,
	}
	tipsy := Dog{
		Name:   "Tipsy",
		Breed:  Poodle,
		Gender: Female,
	}
	_ = bucky
	_ = rocky
	_ = tipsy
}
*/

func (b Breed) String() string {
	switch b {
	case Havanese:
		return "Havanese"
	case Poodle:
		return "Poodle"
	default:
		return "wrong breed"
	}
}

func (g Gender) String() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Wrong gender"
	}
}

// DogSpawner is partial function that help creat Dog with breed and gender
// then return a function that take the name for the Dog struct.
func DogSpawner(breed Breed, gender Gender) NameToDogFunc {
	return func(n Name) Dog {
		return Dog{
			Breed:  breed,
			Gender: gender,
			Name:   n,
		}
	}
}

// function in var
var (
	maleHavaneseSpawner = DogSpawner(Havanese, Male) // return func(n Name) Dog
	femalePoodleSpawner = DogSpawner(Poodle, Female) // return func(n Name) Dog
)

func main() {
	bucky := maleHavaneseSpawner("Bucky")
	rocky := maleHavaneseSpawner("Rocky")
	tipsy := femalePoodleSpawner("Tipsy")

	fmt.Printf("Name: %s, Breed: %s, Gender: %s\n", bucky.Name, bucky.Breed.String(), bucky.Gender.String())
	fmt.Printf("Name: %s, Breed: %s, Gender: %s\n", rocky.Name, rocky.Breed.String(), rocky.Gender.String())
	fmt.Printf("Name: %s, Breed: %s, Gender: %s\n", tipsy.Name, tipsy.Breed.String(), tipsy.Gender.String())
}
