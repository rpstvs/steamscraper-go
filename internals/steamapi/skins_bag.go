package steamapi

import (
	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
)

type Bag struct {
	ID         uuid.UUID
	Name       string
	Items      []uuid.UUID
	TotalPrice float64
}

func CreateBag(name string) Bag {
	var emptySlice []uuid.UUID
	totalPrice := 0.0
	newBag := Bag{
		ID:         uuid.New(),
		Name:       name,
		Items:      emptySlice,
		TotalPrice: totalPrice,
	}
	return newBag

}

func AddItemtoBag(bag database.Bag, item database.GetLatestPriceRow) Bag {
	bag.ItemID = append(bag.ItemID, item.ItemID)

	totalPrice := addPrice(item.Price, bag.Totalvalue)
	updatedBag := &Bag{
		ID:         bag.ID,
		Items:      bag.ItemID,
		TotalPrice: totalPrice,
	}

	return *updatedBag

}

func RemoveItemFromBag(bag database.Bag, item database.GetLatestPriceRow) Bag {

	return Bag{}
}

func addPrice(itemPrice, total float64) float64 {

	total += itemPrice
	return total
}

/*

1- add item to the bag
	- check if the bag exists
	- proceed to add the item
	- add the price to the total of the bag.
*/
