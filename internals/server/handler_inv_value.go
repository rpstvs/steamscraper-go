package server

import (
	"encoding/json"
	"log"
	"net/http"
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

	//inv := steamapi.GetInventory(input.Steamid)
	/*
	   sum := 0

	   	for _, item := range inv.Descriptions {
	   		lastPrice, err := cfg.DB.GetLatestPrice(r.Context(),)
	   	}
	*/
}
