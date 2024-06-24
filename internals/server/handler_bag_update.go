package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

func (cfg *Server) AddItemBag(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		SkinName string    `json:"skinName"`
		IdBag    uuid.UUID `json:"idbag"`
	}
	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(params)

	if err != nil {
		fmt.Println("couldnt decode parameters")
		return
	}

	bag, err := cfg.DB.GetBagbyID(r.Context(), params.IdBag)

	if err != nil {
		fmt.Println("no bag found")
		return
	}

	id, err := cfg.DB.GetItemByName(r.Context(), params.SkinName)

	if err != nil {
		fmt.Println("couldnt find the skin")
	}

	itemPrice, err := cfg.DB.GetLatestPrice(context.Background(), id)

	if err != nil {
		fmt.Println("couldnt retrive the latest price")
		return
	}

	updatedBag := steamapi.AddItemtoBag(bag, itemPrice)

	cfg.DB.UpdateBag(context.Background(), database.UpdateBagParams{
		ID:         updatedBag.ID,
		ItemID:     updatedBag.Items,
		Totalvalue: updatedBag.TotalPrice,
	})

	RespondWithJson(w, http.StatusOK, updatedBag)

}

func (cfg *Server) RemoveItemBag(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		SkinName string    `json:"skinName"`
		IdBag    uuid.UUID `json:"idbag"`
	}
	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(params)

	if err != nil {
		fmt.Println("couldnt decode parameters")
		return
	}

	bag, err := cfg.DB.GetBagbyID(r.Context(), params.IdBag)

	if err != nil {
		fmt.Println("no bag found")
		return
	}

	id, err := cfg.DB.GetItemByName(r.Context(), params.SkinName)

	if err != nil {
		fmt.Println("couldnt find the skin")
	}

	itemPrice, err := cfg.DB.GetLatestPrice(context.Background(), id)

	if err != nil {
		fmt.Println("couldnt retrive the latest price")
		return
	}

	updatedBag := steamapi.AddItemtoBag(bag, itemPrice)

	cfg.DB.UpdateBag(context.Background(), database.UpdateBagParams{
		ID:         updatedBag.ID,
		ItemID:     updatedBag.Items,
		Totalvalue: updatedBag.TotalPrice,
	})

	RespondWithJson(w, http.StatusOK, updatedBag)

}
