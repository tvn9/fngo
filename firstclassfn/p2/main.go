// Learning Type aliases for primities
package main

import "fmt"

// setup type alias veriable of type string
type phoneNumber string
type name string
type age uint
type Person struct {
	name  name
	age   age
	phone phoneNumber
}

// Setter and getter like functions
func (p *Person) setPhoneNumber(s phoneNumber) {
	p.phone = s
}

func (p *Person) getPhoneNumber() phoneNumber {
	return p.phone
}

func (p *Person) setName(s name) {
	p.name = s
}

func (p *Person) getName() name {
	return p.name
}

func (p *Person) upDate(fn name, pn phoneNumber) {
	p.name = fn
	p.phone = pn
}

func (a age) valid() bool {
	return a < 120
}

func isValidPerson(p Person) bool {
	return p.age.valid() && p.name != ""
}

func main() {
	var thanh Person

	thanh.setPhoneNumber("209-888-9999")
	thanh.setName("Kim")

	fmt.Println(thanh.name, thanh.phone)

	thanh.upDate("Thanh Nguyen", "209-555-8888")
	fmt.Println(thanh.name, thanh.phone)

	thanh.age = 53
	fmt.Println("Is thanh's age valide?", thanh.age)

	// Get phone number from person thanh
	fmt.Println(thanh.getPhoneNumber())

	// Get name from thanh
	fmt.Println(thanh.getName())

	// validate age and person
	fmt.Println("Thanh's age is", thanh.age.valid())

	fmt.Println("Is the Thanh a valide persion?", isValidPerson(thanh))

}
