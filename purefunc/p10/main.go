package main

import (
	"fmt"

	"github.com/tvn9/fpgo/purefunc/p10/pkg"
)

func main() {
	myCard := pkg.NewCreditCard(1000)
	hotdog, creditFunc := pkg.OrderHotdog(myCard, pkg.Charge)
	fmt.Printf("%+v\n", hotdog)
	newCard, err := creditFunc()
	if err != nil {
		panic("User has not credit")
	}
	myCard = newCard
	fmt.Printf("%+v\n", myCard)
}
