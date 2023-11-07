// Benchmarking functions
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func immutableCreatePerson() Person {
	p := Person{}
	p = immutableSetName(p, "Thanh")
	p = immutableSetAge(p, 52)
	return p
}

func immutableSetName(p Person, name string) Person {
	p.name = name
	return p
}

func immutableSetAge(p Person, age int) Person {
	p.age = age
	return p
}

func main() {
	p := immutableCreatePerson()
	fmt.Println(p)
}
