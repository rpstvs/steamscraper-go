package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

func (cfg *Server) CreateBag(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name   string    `json:"name"`
		Userid uuid.UUID `json:"userid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}

	err := decoder.Decode(params)

	if err != nil {
		fmt.Println("couldn't decode params")
		return
	}

	bag := steamapi.CreateBag(params.Name)

	bagDb, err := cfg.DB.CreateBag(r.Context(), database.CreateBagParams{
		ID:         bag.ID,
		Totalvalue: 0.0,
		UserID:     params.Userid,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	RespondWithJson(w, http.StatusOK, bagDb)
}
