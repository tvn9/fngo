// DogSpawner exemple using currying function
package main

import "fmt"

type (
	// Name is type alias to be used in Dog object field name
	Name          string
	Breed         byte
	Gender        byte
	NameToDogFunc func(Name) Dog
)

const (
	Havanese Breed = iota
	Poodle
)

const (
	Male Gender = iota
	Female
)

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

func (b Breed) print() string {
	switch b {
	case Havanese:
		return "Havenes"
	case Poodle:
		return "Poodle"
	default:
		return "Not the right breed"
	}
}

func (g Gender) print() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Femail"
	default:
		return "Not exist"
	}
}

func DogSpawner(b Breed, g Gender) NameToDogFunc {
	return func(n Name) Dog {
		return Dog{
			Name:   n,
			Breed:  b,
			Gender: g,
		}
	}
}

func DogSpawnerCurry(b Breed) func(Gender) NameToDogFunc {
	return func(g Gender) NameToDogFunc {
		return func(n Name) Dog {
			return Dog{
				Name:   n,
				Breed:  b,
				Gender: g,
			}
		}
	}
}

func CurriedDogSpawner(b Breed) func(Gender) func(n Name) Dog {
	return func(g Gender) func(Name) Dog {
		return func(n Name) Dog {
			return Dog{
				Name:   n,
				Breed:  b,
				Gender: g,
			}
		}
	}
}

func main() {
	// calling DogSpawner
	dog1 := DogSpawner(Havanese, Male)
	rocky := dog1("Rocky")
	fmt.Printf("Name: %s, Breed: %s, Gender: %s\n", rocky.Name, rocky.Breed.print(), rocky.Gender.print())

	// DogSpawnerCurry
	dog2 := DogSpawnerCurry(Poodle)(Male)
	bucky := dog2("Bucky")
	fmt.Printf("Name: %s, Breed: %s, Gender: %s\n", bucky.Name, bucky.Breed.print(), bucky.Gender.print())

	// CurriedDogSpawner
	dog3 := CurriedDogSpawner(Havanese)(Female)
	tipsy := dog3("Tipsy")
	fmt.Printf("Name: %s, Breed: %s, Gender: %s\n", tipsy.Name, tipsy.Breed.print(), tipsy.Gender.print())
}
