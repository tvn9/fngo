package utils

import "fmt"

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

// NewItemCost
func NewItemCost(sku ItemSku, bPrice, tFee, oFee Amount, item Item) (ItemCost, error) {
	iCost := ItemCost{}
	if item.Sku == sku {
		iCost.Item = item
		iCost.BuyPrice = bPrice
		iCost.TransportFee = tFee
		iCost.OtherFee = oFee
		iCost.Total += (iCost.BuyPrice + iCost.TransportFee + iCost.OtherFee)
		return iCost, nil
	} else {
		return iCost, SKU_MISSMATCH
	}
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

// AddInvItem
func NewItemInv(count ItemCount, item Item, ic ItemCost) (ItemInventory, error) {
	inv := ItemInventory{}
	fmt.Println(count)
	if count > 0 {
		inv.ItemIn = 0
		inv.ItemOnHand = 0
		for i := 1; i <= int(count); i++ {
			inv.Items = append(inv.Items, item)
			inv.ItemIn = ItemCount(len(inv.Items))
			inv.Total += ic.Total
		}
		inv.ItemOnHand = inv.ItemIn
		return inv, nil
	} else {
		return inv, ITEM_COUNT_ERROR
	}
}

// NewInventory
func UpdateInv(item Item, in ItemCount, inv ItemInventory, ic ItemCost) (ItemInventory, error) {
	fmt.Println(inv.Items)
	if len(inv.Items) > 0 {
		inv.ItemIn = ItemCount(0)
		itemOnHand := inv.ItemOnHand

		for i := ItemCount(0); i <= in; i++ {
			inv.Items = append(inv.Items, item)
			inv.ItemIn += ItemCount(i)
		}
		itemOnHand += inv.ItemIn
		inv.ItemOnHand = itemOnHand
		inv.Total += ic.Total
		return inv, nil
	} else {
		return inv, SKU_MISSMATCH
	}
}

// NewInventoryBook
func NewInventoryBook(item Item, itemInv ItemInventory) (InventoryBook, error) {
	newInvBook := InventoryBook{}
	if len(itemInv.Items) <= 0 {
		newInvBook.Items = append(newInvBook.Items, itemInv)
		newInvBook.ItemOnHand = 1
		return newInvBook, nil
	} else {
		return newInvBook, SKU_MISSMATCH
	}
}
