package steamapi

import (
	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
)

type Bag struct {
	id         uuid.UUID
	items      []database.Item
	TotalPrice float64
}

func (c *Client) CreateBag() Bag{

}

func (c *Client) AddItemtoBag(bag Bag, item database.Item)
{
	append(bag, item.ID)
	
	
}

func (c *Client) RemoveItemFromBag() {

}

func addPrice(item database.Price, total float64) float64 {

	total += item.Price
	return total
}

/*

1- add item to the bag
	- check if the bag exists
	- proceed to add the item
	- add the price to the total of the bag.
*/
