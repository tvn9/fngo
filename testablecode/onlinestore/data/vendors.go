
// ShipType holds shipping service type and fee
type ShipType struct {
	ServiceType ShipType
	ShipPrice   ShipFree
}

type Carrier struct {
	CarrierName Carrier
	Phone       Phone
	Email       Email
}

// Shipping service type
type ShipService struct {
	ShipType []ShipType
	Carrier  []Carrier
	ShipFree ShipFee
}

// Shipping cart
type ShipCart struct {
	ShipType ShipService
	Carrier  Carrier
	ShipFee  ShipFee
}

type CreditCard struct {
	credit int
}

type CheckOutCart struct {
	Item    []Item
	ShipFee ShipCart
	Tax
	CardTotal int
}

type Vendor struct {
	CompanyName CompanyName
	Address     Address
	Phone       Phone
	Email       Email
}

// ShipType holds shipping service type and fee
type ShipType struct {
	ServiceType ShipType
	ShipPrice   ShipFree
}

type Carrier struct {
	CarrierName Carrier
	Phone       Phone
	Email       Email
}

// Shipping service type
type ShipService struct {
	ShipType []ShipType
	Carrier  []Carrier
	ShipFree ShipFee
}

type CheckOutCart struct {
	Item    []Item
	ShipFee ShipCart
	Tax
	CardTotal int
}
