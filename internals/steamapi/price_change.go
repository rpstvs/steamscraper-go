package steamapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
)

func (cfg *Client) PriceChangeDaily(itemname string) {
	// Pegar nos dois ultimos preços de um item
	// calculcar a change
	// escrever a change para o item

	id, _ := cfg.DB.GetItemByName(context.Background(), itemname)

	item,_ := cfg.DB.GetItemRecord(context.Background(), database.GetItemRecordParams{
		ItemID: id,
		Limit:  2,
	})
	
	dailyChange := (item[0] + item[1])/2

	cfg.DB.UpdateDailyChange(context.Background(), database.UpdateDailyChangeParams{
		Daychange: dailyChange,
		ID: id,
	})


	
}

func (cfg *Client) WeeklyPriceChange
