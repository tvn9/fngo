package utils

type Item struct {
	Sku         ItemSku
	Name        ItemName
	Description ItemDesc
	Active      bool
}

type ItemCost struct {
	Item         Item
	BuyPrice     Amount
	TransportFee Amount
	OtherFee     Amount
	Total        Amount
}

type PriceBook struct {
}

type ItemInventory struct {
	Item       Item
	ItemIn     ItemCount
	ItemOut    ItemCount
	ItemOnHand ItemCount
	Total      Amount
}

type InventoryAll struct {
	Items         []ItemInventory
	InventoryCost Amount
}

// NewItem creates a new item
