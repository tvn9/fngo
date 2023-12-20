package items

import (
	"onlinestore/utils/skus"
	"testing"
)

func TestNewItem(t *testing.T) {
	sku := "tn0001"
	newSku, _ := skus.NewSku(sku, skus.SkuFormatCheck)
	newItem, err := NewItem(newSku)

	if err != nil || newItem.Sku != newSku {
		t.Errorf("expected %v, but go %v, %v", newSku, newItem.Sku, err)
	} else {
		t.Logf("test pass %v = %v, error: %v", newItem, newSku, err)
	}
}

func TestAddItemName(t *testing.T) {
	sku := "tn0002"
	iname := "Men Hat Blue"
	newSku, _ := skus.NewSku(sku, skus.SkuFormatCheck)
	newItem, _ := NewItem(newSku)
	iName := ItemName(iname)

	nItem, err := AddItemName(newSku, iName, newItem)

	if nItem.Sku != newItem.Sku || nItem.Name != iName || err != nil {
		t.Errorf("expected %v, but got %v, Error %v", iName, nItem.Name, err)
	}
}
