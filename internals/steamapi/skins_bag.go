package steamapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
)

type Bag struct {
	id         uuid.UUID
	items      []database.Item
	TotalPrice float64
}

func (c *Client) CreateBag() Bag {
	var emptySlice []database.Item
	totalPrice := 0.0
	newBag := &Bag{
		id:         uuid.New(),
		items:      emptySlice,
		TotalPrice: totalPrice,
	}
	return *newBag

}

func (c *Client) AddItemtoBag(bag Bag, item database.Item) Bag {
	bag.items = append(bag.items, item)

	itemPrice, err := c.DB.GetLatestPrice(context.Background(), item.ID)

	if err != nil {
		fmt.Println("item does not exist")
		return Bag{}
	}

	totalPrice := addPrice(itemPrice.Price, bag.TotalPrice)
	updatedBag := &Bag{
		id:         bag.id,
		items:      bag.items,
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
