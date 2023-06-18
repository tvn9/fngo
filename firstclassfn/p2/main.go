// Learning Type aliases for primities
package main

import "fmt"

// phoneNumber is the type alias veriable of type string
type phoneNumber string
type name string
type age uint
type Person struct {
	name  name
	age   age
	phone phoneNumber
}

func (p *Person) setPhoneNumber(s phoneNumber) {
	p.phone = s
}

func (p *Person) setName(s name) {
	p.name = s
}

func (p *Person) upDate(fn name, pn phoneNumber) {
	p.name = fn
	p.phone = pn
}

func main() {
	var thanh Person

	thanh.setPhoneNumber("209-888-9999")
	thanh.setName("Kim")

	fmt.Println(thanh.name, thanh.phone)

	thanh.upDate("Thanh Nguyen", "209-555-8888")
	fmt.Println(thanh.name, thanh.phone)

	thanh.age = 53
	fmt.Println("Thanh's age is", thanh.age)
}
