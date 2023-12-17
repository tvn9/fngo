package utils

import "errors"

// type alias for general standard name
type (
	FirstName   string
	LastName    string
	MidName     string
	SSN         string
	DriverID    string
	CompanyName string
	Phone       string
	Email       string
	AddrNum     string
	Street      string
	City        string
	State       string
	ZipCode     string
	Country     string
)

// type alias for calculation
type (
	Currency string
	Amount   uint
)

// type alias for Item struct
type (
	ItemSku   string
	ItemName  string
	ItemDesc  string
	ItemPrice uint
	ItemCount uint
)

// Type alias for Cart struct
type (
	TotalItem uint
	Subtotal  int
	Tax       uint
	ShipFree  uint
	Total     int
)

// Type alias for shipping
type (
	ShipType string
	Carrier  string
	ShipFee  uint
)

// Error
type ItemError error

// Error types
var (
	SKU_MISSMATCH ItemError = ItemError(errors.New("sku does not macth the record"))
)
