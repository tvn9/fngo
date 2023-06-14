package main

import "fmt"

type phoneNumber string
type Person struct {
	name        string
	phonenumber phoneNumber
}

func (p *Person) setPhoneNumber(s phoneNumber) {
	p.phonenumber = s
}

func (p *Person) setName(s string) {
	p.name = s
}

func main() {

	thanh := Person{}

	thanh.setPhoneNumber("209-888-9999")
	thanh.setName("Thanh Nguyen")

	fmt.Println(thanh.name, thanh.phonenumber)
}
