package steamapi

import (
	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
)

type Bag struct {
	id    uuid.UUID
	items []database.Item
}

func (c *Client) AddItemtoBag() {

}

func (c *Client) RemoveItemFromBag() {

}

/*

1- add item to the bag
	- check if the bag exists
	- proceed to add the item
	- add the price to the total of the bag.
*/
