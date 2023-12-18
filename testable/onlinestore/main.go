package main

import (
	"fmt"
	"fpgo/testable/onlinestore/utils"
	"log"
)

func main() {
	// Create new Person
	fn := utils.FirstName("Thanh")
	thanh := utils.NewPerson(fn)
	fmt.Printf("New Person: %v\n", thanh)

	// Update person
	ls := utils.LastName("Nguyen")
	uThanh := utils.UpdatePersonDetails(thanh.FirstName, ls, thanh)
	fmt.Printf("Update Person: %v\n", uThanh)

	// Create new contact
	cThanh := utils.NewContact(uThanh)
	fmt.Printf("NewContact: %v\n", cThanh)

	// Create new contact
	addrNum := utils.AddrNum("1001")
	str1001 := utils.Street("St. Peter")
	tPhone := utils.Phone("209-555-8888")
	tEmail := utils.Email("thanh@email.com")
	addr := utils.NewAddress(addrNum, str1001)
	fmt.Printf("New Address: %v\n", addr)

	// Update address details
	city := utils.City("Tracy")
	state := utils.State("California")
	zip := utils.ZipCode("95376")
	usa := utils.Country("United State")
	taddr := utils.UpdateAddrDetails(addr.AddressNum, addr.Street, city, state, zip, usa, addr)
	fmt.Printf("Update Address: %v\n", taddr)

	// Update existing contact
	ct := utils.UpdateContactDetails(uThanh, taddr, tPhone, tEmail, cThanh)
	fmt.Printf("Update Contact %v\n", ct)

	// Create new item
	sku := utils.ItemSku("TN0001")
	desc := utils.ItemDesc("Standard Size Hotdog with Beef")
	bPrice := utils.Amount(12)
	tFee := utils.Amount(3)
	oFee := utils.Amount(1)

	tn0001 := utils.NewItem(sku)
	fmt.Printf("Create New Item: %v\n", tn0001)

	// UpdateItemDetails
	tn0001, err := utils.UpdateItemDetails(sku, utils.ItemName("HotDog"), desc, tn0001)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Update Item Details: %v\n", tn0001)

	// Set new item cost
	icost, err := utils.NewItemCost(sku, bPrice, tFee, oFee, tn0001)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Item Cost: %v\n", icost)

	// Create new inventory item
	invItem, err := utils.NewItemInv(utils.ItemCount(10), tn0001, icost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Item cost %v\n", icost)
	for i, item := range invItem.Items {
		fmt.Printf("#%d, Sku: %s, Name %s, Desc %s, Count: %d, Cost: %d, Inv Cost: %d\n",
			i, item.Sku, item.Name, item.Description, invItem.ItemOnHand, icost.Total, invItem.Total)
	}

	/*
		// UpdateInv
		uInvItem, err := utils.UpdateInv(tn0001, utils.ItemCount(5), itemInv, icost)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Update Inventory: %v\n", uInvItem)
	*/

}
