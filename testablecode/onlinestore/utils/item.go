package utils

type Item struct {
	Sku         ItemSku
	Name        ItemName
	Description ItemDesc
	Active      bool
}

type ItemCost struct {
	Item         Item
	BuyPrice     ItemPrice
	TransportFee ItemPrice
	OtherFee     ItemPrice
	Total        Amount
}

type ItemInventory struct {
	Items      []Item
	ItemIn     ItemCount
	ItemOut    ItemCount
	ItemOnHand ItemCount
	Total      Amount
}

type InventoryBook struct {
	Items         []ItemInventory
	NumOfItems    ItemCount
	ItemOnHand    ItemCount
	InventoryCost Amount
}

// NewItem creates a new item
func NewItem(sku ItemSku) Item {
	return Item{Sku: sku, Active: true}
}

// UpdateItemDetails
func UpdateItemDetails(sku ItemSku, name ItemName, desc ItemDesc, item Item) (Item, error) {
	if item.Sku == sku {
		item.Name = name
		item.Description = desc
		item.Active = true
		return item, nil
	} else {
		return item, SKU_MISSMATCH
	}
}

// NewItemCost
func NewItemCost(sku ItemSku, bPrice, tFee, oFee ItemPrice, item Item) (ItemCost, error) {
	iCost := ItemCost{}
	if item.Sku == sku {
		iCost.Item = item
		iCost.BuyPrice = bPrice
		iCost.TransportFee = tFee
		iCost.OtherFee = oFee
		iCost.Total += Amount(iCost.BuyPrice + iCost.TransportFee + iCost.OtherFee)
		return iCost, nil
	} else {
		return iCost, SKU_MISSMATCH
	}
}

// NewInventory
func NewInventory(sku ItemSku, in ItemCount, item Item, ic ItemCost) (ItemInventory, error) {
	newInv := ItemInventory{}
	if item.Sku == sku {
		inCount := in
		newInv.ItemIn = ItemCount(0)
		for i := ItemCount(1); i <= inCount; i++ {
			newInv.Items = append(newInv.Items, item)
			newInv.ItemIn += ItemCount(i)
		}
		newInv.ItemOnHand += newInv.ItemIn
		newInv.Total += ic.Total
		return newInv, nil
	} else {
		return newInv, SKU_MISSMATCH
	}
}

// NewInventoryBook
func NewInventoryBook(sku ItemSku, item Item, itemInv ItemInventory) (InventoryBook, error) {
	newInvBook := InventoryBook{}
	if item.Sku == sku {
		newInvBook.Items = append(newInvBook.Items, itemInv)
		newInvBook.ItemOnHand += ItemCount(1)
		return newInvBook, nil
	} else {
		return newInvBook, SKU_MISSMATCH
	}
}
