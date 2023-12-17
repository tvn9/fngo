package main

import (
	"fmt"
	"fpgo/testablecode/onlinestore/utils"
)

func main() {
	// Create new Person
	fn := utils.FirstName("Thanh")
	thanh := utils.NewPerson(fn)
	fmt.Printf("New Person: %v\n", thanh)

	// Update person
	ls := utils.LastName("Nguyen")
	uThanh := utils.UpdatePersonDetails(thanh.FirstName, ls, thanh)
	fmt.Printf("Update Person: %v\n", uThanh)

	// Create new contact
	cThanh := utils.NewContact(uThanh)
	fmt.Printf("NewContact: %v\n", cThanh)

	// Create new contact
	addrNum := utils.AddrNum("1001")
	str1001 := utils.Street("St. Peter")
	tPhone := utils.Phone("209-555-8888")
	tEmail := utils.Email("thanh@email.com")
	addr := utils.NewAddress(addrNum, str1001)
	fmt.Printf("New Address: %v\n", addr)

	// Update address details
	city := utils.City("Tracy")
	state := utils.State("California")
	zip := utils.ZipCode("95376")
	usa := utils.Country("United State")
	taddr := utils.UpdateAddrDetails(addr.AddressNum, addr.Street, city, state, zip, usa, addr)
	fmt.Printf("Update Address: %v\n", taddr)

	// Update existing contact
	ct := utils.UpdateContactDetails(uThanh, taddr, tPhone, tEmail, cThanh)
	fmt.Printf("Update Contact %v\n", ct)
}
