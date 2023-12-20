package main

import (
	"fmt"
	"log"
	"onlinestore/types"
	"onlinestore/utils/items"
	"onlinestore/utils/skus"
)

func main() {
	sks := []string{"tn0001", "tn0002", "tn0003", "tn0004"}

	newSkusMap := make(map[types.Sku]types.Sku)
	for _, s := range sks {
		testStr := skus.SkuFormatCheck(s)
		fmt.Println(testStr)

		sku, err := skus.NewSku(s, skus.SkuFormatCheck)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(sku)

		skus.UpdateSkusMap(sku, newSkusMap)
	}
	fmt.Println(newSkusMap, len(newSkusMap))

	// Generate Sku
	str := "tn0001"
	newSkuList := skus.GenerateSkus(str, 5)
	fmt.Println(newSkuList)

	// Create new Sku
	newSku, _ := skus.NewSku(str, skus.SkuFormatCheck)
	fmt.Println("New Sku: ", newSku)

	newItem, err := items.NewItem(newSku)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create New Item: ", newItem)

	// Add ItemName
	newIName := items.ItemName("Men Hat Blue")
	item, err := items.AddItemName(newSku, newIName, newItem)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(item)
}
