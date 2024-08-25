package server

/*
type response struct {
	Bagid      uuid.UUID `json:"bagid"`
	Totalvalue float64   `json:"totalvalue"`
	Items      struct {
		Itemid uuid.UUID `json:"itemid"`
		Amount int32     `json:"amount"`
	} `json:"items"`
}

func (cfg *Server) AddItemBag(w http.ResponseWriter, r *http.Request, user database.User) {
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

	itemPrice, err := cfg.DB.GetLatestPrice(r.Context(), id)

	if err != nil {
		fmt.Println("couldnt retrive the latest price")
		return
	}

	amountInBag, err := cfg.DB.GetBagItem(r.Context(), database.GetBagItemParams{
		BagID:  bag.ID,
		ItemID: id,
	})

	if amountInBag == 0 {
		cfg.DB.AddItem(r.Context(), database.AddItemParams{
			BagID:  bag.ID,
			ItemID: id,
			Amount: params.Amount,
		})

		newTotal, _ := steamapi.AddPrice(itemPrice.Price, bag.Totalvalue, params.Amount, amountInBag)

		cfg.DB.UpdateBag(r.Context(), database.UpdateBagParams{
			ID:         bag.ID,
			Totalvalue: newTotal,
		})

		return
	}

	newTotal, newAmount := steamapi.AddPrice(itemPrice.Price, bag.Totalvalue, params.Amount, amountInBag)

	cfg.DB.UpdateBagItem(r.Context(), database.UpdateBagItemParams{
		BagID:  bag.ID,
		ItemID: id,
		Amount: newAmount,
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
			Amount: newAmount,
		},
	})

}
*/
