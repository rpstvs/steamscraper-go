package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func (cfg *ApiConfig) HandlerGetPric(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Skin string `json:"skin"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		fmt.Println("couldn't decode params")
	}

	id, err := cfg.DB.GetItemByName(r.Context(), params.Skin)

	if err != nil {
		fmt.Println("couldnt find the skin")
	}

	price := cfg.DB.

}
