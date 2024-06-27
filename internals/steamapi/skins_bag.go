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

func AddItemtoBag(bag database.BagItem, item database.GetLatestPriceRow) {

}

func RemoveItemFromBag(bag database.Bag, item database.GetLatestPriceRow) Bag {

	return Bag{}
}

func AddPrice(itemPrice, total float64, amount int32) float64 {

	total += itemPrice * float64(amount)
	return total
}

func SubPrice(itemPrice, total float64, amount int32) float64 {

	total -= itemPrice
	return total
}
