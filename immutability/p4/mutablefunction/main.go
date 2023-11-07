// Benchmarking functions
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func mutableCreatePerson() *Person {
	p := &Person{}
	mutableSetName(p, "Thanh")
	mutableSetAge(p, 52)
	return p
}

func mutableSetName(p *Person, name string) {
	p.name = name
}

func mutableSetAge(p *Person, age int) {
	p.age = age
}

func main() {
	person := mutableCreatePerson()
	fmt.Println(person)
}
