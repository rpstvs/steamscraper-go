package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

func (cfg *Server) CreateBag(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}

	err := decoder.Decode(params)

	if err != nil {
		fmt.Println("couldn't decode params")
		return
	}

	bag := steamapi.CreateBag(params.Name)

	cfg.DB.CreateBag(context.Background(), database.CreateBagParams{
		ID:         bag.ID,
		ItemID:     bag.Items,
		Totalvalue: bag.TotalPrice,
	})

	RespondWithJson(w, http.StatusOK, bag)
}
