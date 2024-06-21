package steamapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
)

type Bag struct {
	ID         uuid.UUID
	Name       string
	Items      []database.Item
	TotalPrice float64
}

func CreateBag(name string) Bag {
	var emptySlice []database.Item
	totalPrice := 0.0
	newBag := Bag{
		ID:         uuid.New(),
		Name:       name,
		Items:      emptySlice,
		TotalPrice: totalPrice,
	}
	return newBag

}

func (c *Client) AddItemtoBag(bag Bag, item database.Item) Bag {
	bag.Items = append(bag.Items, item)

	itemPrice, err := c.DB.GetLatestPrice(context.Background(), item.ID)

	if err != nil {
		fmt.Println("item does not exist")
		return Bag{}
	}

	totalPrice := addPrice(itemPrice.Price, bag.TotalPrice)
	updatedBag := &Bag{
		ID:         bag.ID,
		Items:      bag.Items,
		TotalPrice: totalPrice,
	}

	return *updatedBag

}

func (c *Client) RemoveItemFromBag() {

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
