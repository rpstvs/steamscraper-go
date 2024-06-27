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

type response struct {
	Bagid      uuid.UUID `json:"bagid"`
	Totalvalue float64   `json:"totalvalue"`
	Items      struct {
		Itemid uuid.UUID `json:"itemid"`
		Amount int32     `json:"amount"`
	} `json:"items"`
}

func (cfg *Server) AddItemBag(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		SkinName string    `json:"skinName"`
		IdBag    uuid.UUID `json:"idbag"`
		Amount   int32     `json:"amount"`
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

	newTotal := steamapi.AddPrice(itemPrice.Price, bag.Totalvalue, params.Amount)

	cfg.DB.AddItem(r.Context(), database.AddItemParams{
		BagID:  bag.ID,
		ItemID: id,
		Amount: params.Amount,
	})

	cfg.DB.UpdateBag(r.Context(), database.UpdateBagParams{
		ID:         bag.ID,
		Totalvalue: newTotal,
	})

	RespondWithJson(w, http.StatusOK, response{
		Bagid:      bag.ID,
		Totalvalue: newTotal,
		Items: struct {
			Itemid uuid.UUID "json:\"itemid\""
			Amount int32     "json:\"amount\""
		}{
			Itemid: id,
			Amount: params.Amount,
		},
	})

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

	cfg.DB.UpdateBag(context.Background(), database.UpdateBagParams{
		ID:         bag.ID,
		Totalvalue: bag.Totalvalue,
	})

	RespondWithJson(w, http.StatusOK, response{
		Bagid:      bag.ID,
		Totalvalue: itemPrice.Price,
		Items: struct {
			Itemid uuid.UUID "json:\"itemid\""
			Amount int32     "json:\"amount\""
		}{
			Itemid: id,
			Amount: 3,
		},
	})

}
