package main

import "fmt"

type Person struct {
	name string
	age  int
}

// setName sets the person name for Person struct
// however, this function signature does not change the name
// because p is just a copy of Person struct.
func setName(p Person, name string) {
	p.name = name
}

// setPersonName updates the Person's name by pointer
func setPersonName(p *Person, name string) {
	p.name = name
}

func setNameReturnPerson(p Person, name string) Person {
	p.name = name
	return p
}

func main() {
	p := Person{
		name: "Benny",
		age:  55,
	}
	setName(p, "Bjorn")
	fmt.Println(p.name)

	setPersonName(&p, "Bjorn")
	fmt.Println(p.name)

	person := setNameReturnPerson(p, "Thanh")
	fmt.Println(person)

}
