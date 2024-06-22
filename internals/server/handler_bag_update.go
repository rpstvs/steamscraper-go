package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

}
