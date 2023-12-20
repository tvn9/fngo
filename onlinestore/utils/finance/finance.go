package finance

import "onlinestore/types"

// type alias for finance calculation
type Amount uint

// ItemCost holds the total inventory cost for each new created item
type ItemCost struct {
	Sku          types.Sku
	BuyPrice     Amount
	TransportFee Amount
	OtherFee     Amount
	Total        Amount
}
