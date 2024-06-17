package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (cfg *Server) GetLatestPrice(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}

	fmt.Println("got a request")
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		fmt.Println("couldn't decode params")
		return
	}

	id, err := cfg.DB.GetItemByName(r.Context(), params.Body)

	if err != nil {
		fmt.Println("couldnt find the skin")
	}

	price, err := cfg.DB.GetPricebyId(r.Context(), id)

	if err != nil {
		fmt.Println("no price for this skin")
		return
	}

	item := &Item{
		Name:      params.Body,
		Price:     price.Price,
		PriceDate: price.Pricedate,
	}
	fmt.Println(item)
	RespondWithJson(w, http.StatusOK, item)

}
