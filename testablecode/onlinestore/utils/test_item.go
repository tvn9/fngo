package utils

import (
	"testing"
)

func TestNewItem(t *testing.T) {

	sku := ItemSku("TN001")
	newItem := NewItem(sku)

	t.Errorf("expected %v, but got %v", sku, newItem.Sku)

}
