package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/steamapi"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func (cfg *Server) inventoryValue(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Steamid string `json:"steamid"`
	}

	type response struct {
		Value float64 `json:"value"`
		Items int     `json:"items"`
	}

	var input parameters
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&input)

	if err != nil {
		log.Println("couldnt decode steam id")
	}

	fmt.Println(input.Steamid)

	inv := steamapi.GetInventory(input.Steamid)

	var itemClassids []string

	quantity := utils.QuantityMap(inv)

	for k, _ := range quantity {
		itemClassids = append(itemClassids, k)
	}

	resp, err := cfg.DB.GetBatchPrices(r.Context(), itemClassids)

	if err != nil {
		log.Println("Couldnt get the batch of items")
	}
	var sum float64
	totalItems := 0
	for _, v := range resp {
		sum += v.PricesPrice * float64(quantity[v.Classid])
		totalItems += quantity[v.Classid]
		fmt.Printf("adding item %s with price %f, with amount %d, to the sum: %f\n", v.Itemname, v.PricesPrice, quantity[v.Classid], sum)
	}

	RespondWithJson(w, 200, response{
		Value: sum,
		Items: totalItems,
	})

}
