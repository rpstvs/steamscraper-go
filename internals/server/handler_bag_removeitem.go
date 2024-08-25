package server

/*
func (cfg *Server) RemoveItemBag(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		SkinName string    `json:"skinName"`
		IdBag    uuid.UUID `json:"idbag"`
		Amount   int32     `json:"Amount"`
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

	amountInBag, err := cfg.DB.GetBagItem(r.Context(), database.GetBagItemParams{
		BagID:  bag.ID,
		ItemID: id,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	newTotal, newAmount := steamapi.SubPrice(itemPrice.Price, bag.Totalvalue, amountInBag, params.Amount)

	if newAmount < 0 {
		fmt.Print("no items left in the bag")
		return
	}

	if newAmount == 0 {
		cfg.DB.UpdateBag(r.Context(), database.UpdateBagParams{
			ID:         bag.ID,
			Totalvalue: newTotal,
		})

		cfg.DB.DeleteItem(r.Context(), database.DeleteItemParams{
			BagID:  bag.ID,
			ItemID: id,
		})
		fmt.Println("item apagado.-")
		return
	}

	cfg.DB.UpdateBag(r.Context(), database.UpdateBagParams{
		ID:         bag.ID,
		Totalvalue: newTotal,
	})

	cfg.DB.UpdateBagItem(r.Context(), database.UpdateBagItemParams{
		BagID:  bag.ID,
		Amount: newAmount,
		ItemID: id,
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
