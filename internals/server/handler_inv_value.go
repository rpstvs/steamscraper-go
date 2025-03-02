package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/steamapi"
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

	inv := steamapi.GetInventory(input.Steamid)

	sum := 0

	for _, item := range inv.Descriptions {
		lastPrice, err := cfg.DB.GetPrice(r.Context(), item.MarketHashName)
		if err != nil {
			fmt.Println(err)
			continue

		}
		sum += int(lastPrice)
		fmt.Printf("adding the price %f, of item %s", lastPrice, item.MarketHashName)
	}

	fmt.Printf("The sum is: %d \n", sum)

}
