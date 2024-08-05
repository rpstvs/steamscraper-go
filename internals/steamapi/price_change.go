package steamapi

import (
	"context"
	"fmt"

	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func (cfg *Client) PriceChangeDaily(itemname string) {

	ctx := context.Background()
	id, _ := cfg.DB.GetItemByName(ctx, itemname)

	item, _ := cfg.DB.GetItemRecord(ctx, database.GetItemRecordParams{
		ItemID: id,
		Limit:  2,
	})

	if len(item) < 2 {
		cfg.DB.UpdateDailyChange(ctx, database.UpdateDailyChangeParams{
			ID:        id,
			Daychange: 0.01,
		})
		return
	}

	dailyChange := utils.DailyPriceChange(item[0], item[1])
	fmt.Println(item[1], item[0], dailyChange)

	cfg.DB.UpdateDailyChange(ctx, database.UpdateDailyChangeParams{
		Daychange: dailyChange,
		ID:        id,
	})

}

func (cfg *Client) WeeklyPriceChange(itemname string) {

	ctx := context.Background()

	id, _ := cfg.DB.GetItemByName(ctx, itemname)

	item, _ := cfg.DB.GetItemRecord(ctx, database.GetItemRecordParams{
		ItemID: id,
		Limit:  7,
	})

	weeklyChage := utils.WeeklyPriceChange(item)
	cfg.DB.UpdateWeeklyChange(ctx, database.UpdateWeeklyChangeParams{
		ID:         id,
		Weekchange: weeklyChage,
	})

}
