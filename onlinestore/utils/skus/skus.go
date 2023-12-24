package skus

import (
	"errors"
	"onlinestore/types"
	"strconv"
	"strings"
)

// Error types
type skuError error

type Skus []types.Sku

// Sku error type
var (
	NOT_VALID_SKU_FORMAT skuError = skuError(errors.New("not a valid Sku format"))
)

// skuMap
type SkusMap map[types.Sku]types.Sku

// NewSku create a new Sku object that hold on the Sku number
func NewSku(s string, valid types.ValidString) (types.Sku, error) {
	skuNum := strings.ToUpper(s)
	newSku := types.Sku("")
	if valid(skuNum) {
		newSku = types.Sku(skuNum)
		return newSku, nil
	} else {
		return newSku, NOT_VALID_SKU_FORMAT
	}
}

// NewSkusMap creates a map
func NewSkusMap(s types.Sku) SkusMap {
	sm := make(map[types.Sku]types.Sku)
	sm[s] = s
	return sm
}

// Sku format check
func SkuFormatCheck(s string) bool {
	us := strings.ToUpper(s)
	str := us
	// Pretend this is a validation for legal Sku
	// Future update should include regular expression to check for legal Sku format
	if str == us {
		return true
	} else {
		return false
	}
}

// UpdateSkusMap
func UpdateSkusMap(s types.Sku, sm SkusMap) SkusMap {
	sm[s] = s
	return sm
}

// GenerateSkus
func GenerateSkus(sku string, n int) Skus {
	usku := strings.ToUpper(sku)
	startNum, _ := strconv.Atoi(string(usku[len(usku)-1]))
	skus := Skus{}
	if SkuFormatCheck(usku) {
		for i := 1; i <= n; i++ {
			str := strconv.Itoa(startNum + i)
			skus = append(skus, types.Sku(usku[:5]+str))
		}
	}
	return skus
}
