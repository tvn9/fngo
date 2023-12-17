package utils

// Address is address object
type Address struct {
	AddressNum AddrNum
	Street     Street
	City       City
	State      State
	Zip        ZipCode
	Country    Country
}

// NewAddress creates a partial address with just address number and street name
func NewAddress(addr AddrNum, street Street) Address {
	return Address{AddressNum: addr, Street: street}
}

// UpdateAddrDetails
func UpdateAddrDetails(
	addrNum AddrNum,
	street Street,
	city City,
	state State,
	zip ZipCode,
	country Country,
	addr Address) Address {
	addr.AddressNum = addrNum
	addr.Street = street
	addr.City = city
	addr.State = state
	addr.Country = country
	return addr
}
