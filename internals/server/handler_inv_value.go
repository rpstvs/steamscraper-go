package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/steamapi"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

//get inventory through steamid

func (cfg *Server) inventoryValue(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Steamid string `json:"steamid"`
	}

	var input parameters
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&input)

	if err != nil {
		log.Println("couldnt decode steam id")
	}

	fmt.Println(input.Steamid)

	var pages []utils.Inventory

	steamapi.GetInventory(input.Steamid, "0", &pages)

	var itemNames []string

	for _, page := range pages {

		for _, item := range page.Descriptions {
			itemNames = append(itemNames, item.MarketHashName)
		}

	}

	resp, err := cfg.DB.GetBatchPrices(r.Context(), itemNames)

	if err != nil {
		log.Println("Couldnt get the batch of items")
	}

	itemsMap := make(map[string]float64, len(resp))

	for _, item := range resp {
		itemsMap[item.Itemname] = item.PricesPrice
	}
	var sum float64
	for _, v := range itemsMap {
		sum += v
	}

	fmt.Printf("You have %d items and the value is %f.2 \n", len(itemNames), sum)

}
