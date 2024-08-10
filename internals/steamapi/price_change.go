package steamapi

import (
	"context"

	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func (cfg *Client) PriceChangeDaily(itemname string) {

	ctx := context.Background()
	id, err := cfg.DB.GetItemByName(ctx, itemname)

	if err != nil {
		return
	}

	item, _ := cfg.DB.GetItemRecord(ctx, database.GetItemRecordParams{
		ItemID: id,
		Limit:  2,
	})

	if len(item) < 2 {
		cfg.DB.UpdateDailyChange(ctx, database.UpdateDailyChangeParams{
			ID:        id,
			Daychange: 0.00,
		})
		return
	}

	dailyChange := utils.DailyPriceChange(item[0], item[1])

	cfg.DB.UpdateDailyChange(ctx, database.UpdateDailyChangeParams{
		Daychange: dailyChange,
		ID:        id,
	})

}

func (cfg *Client) WeeklyPriceChange(itemname string) {

	ctx := context.Background()

	id, err := cfg.DB.GetItemByName(ctx, itemname)

	if err != nil {
		return
	}

	item, err := cfg.DB.GetItemRecord(ctx, database.GetItemRecordParams{
		ItemID: id,
		Limit:  7,
	})

	if err != nil {
		cfg.DB.UpdateWeeklyChange(ctx, database.UpdateWeeklyChangeParams{
			ID:         id,
			Weekchange: 0.00,
		})
		return
	}

	weeklyChage := utils.WeeklyPriceChange(item)
	cfg.DB.UpdateWeeklyChange(ctx, database.UpdateWeeklyChangeParams{
		ID:         id,
		Weekchange: weeklyChage,
	})

}
