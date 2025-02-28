package server

import (
	"encoding/json"
	"log"
	"net/http"
)

//get inventory through steamid

func (cfg *Server) inventoryValue(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	var input parameters
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&input)

	if err != nil {
		log.Println("couldnt decode steam id")
	}

	const baseurl = "https://steamcommunity.com/inventory/{user_id}/730/2"

}
