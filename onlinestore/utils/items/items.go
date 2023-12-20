package items

import (
	"errors"
	"onlinestore/types"
)

// ItemErrors of type error
type ItemError error

// itemErrors
var (
	INVALID_SKU       ItemError = ItemError(errors.New("invalid sku"))
	SKU_DOESNOT_MATCH ItemError = ItemError(errors.New("sku does not match the record"))
)

// type alias for Item struct
type (
	ItemName  string
	ItemDesc  string
	ItemCount uint
	Count     uint
)

// Item hold the item Sku ID
type Item struct {
	Sku    types.Sku
	Name   ItemName
	Desc   ItemDesc
	Active bool
}

// NewItem create a new item
func NewItem(s types.Sku) (Item, error) {
	i := Item{}
	if i.Sku == "" {
		i.Sku = s
		i.Active = true
		return i, nil
	} else {
		return i, INVALID_SKU
	}
}

// AddItemDetails
func AddItemName(s types.Sku, n ItemName, item Item) (Item, error) {
	if s == item.Sku {
		item.Name = n
		return item, nil
	} else {
		return item, SKU_DOESNOT_MATCH
	}
}
