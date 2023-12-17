package utils

// Name
type Name struct {
	FirstName FirstName
	LastName  LastName
}

// Person
type Person struct {
	Name      Name
	FirstName FirstName
	LastName  LastName
}

// NewName
func NewName(fn FirstName) Name {
	return Name{FirstName: fn}
}

// CreatePerson
func NewPerson(fn FirstName) Person {
	fName := NewName(fn)
	return Person{Name: fName, FirstName: fName.FirstName}
}

// Update Person Details
func UpdatePersonDetails(fn FirstName, ls LastName, p Person) Person {
	name := Name{}
	name.FirstName = fn
	name.LastName = ls
	p.Name = name
	p.FirstName = name.FirstName
	p.LastName = name.LastName
	return p
}
