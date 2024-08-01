package steamapi

import (
	"context"

	"github.com/rpstvs/steamscraper-go/internals/database"
)

func (cfg *Client) PriceChangeDaily(itemname string) {
	// Pegar nos dois ultimos preços de um item
	// calculcar a change
	// escrever a change para o item
	ctx := context.Background()
	id, _ := cfg.DB.GetItemByName(ctx, itemname)

	item, _ := cfg.DB.GetItemRecord(ctx, database.GetItemRecordParams{
		ItemID: id,
		Limit:  2,
	})

	dailyChange := (item[0] + item[1]) / 2

	cfg.DB.UpdateDailyChange(ctx, database.UpdateDailyChangeParams{
		Daychange: dailyChange,
		ID:        id,
	})

}

func (cfg *Client) WeeklyPriceChange(itemname string) {

}
