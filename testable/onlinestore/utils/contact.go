package utils

type Contact struct {
	Person  Person
	Address Address
	Phone   Phone
	Email   Email
}

// Create New Contact
func NewContact(p Person) Contact {
	return Contact{Person: p}
}

// UpdateContactDetails
func UpdateContactDetails(person Person, address Address, phone Phone, email Email, contact Contact) Contact {
	contact.Person = person
	contact.Address = address
	contact.Phone = phone
	contact.Email = email
	return contact
}
