package steamapi

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func (cfg *Client) PriceChangeDaily(id uuid.UUID, record []float64, itemname string, ctx context.Context) {

	if len(record) < 2 {
		cfg.DB.UpdateDailyChange(ctx, database.UpdateDailyChangeParams{
			ID:        id,
			Daychange: 0.00,
		})
		return
	}

	dailyChange := utils.DailyPriceChange(record[0], record[1])

	cfg.DB.UpdateDailyChange(ctx, database.UpdateDailyChangeParams{
		Daychange: dailyChange,
		ID:        id,
	})

	log.Printf("Item: %s - Price Today: %f - Old Price: %f -  DayChange: %f\n", itemname, record[0], record[1], dailyChange)

}

func (cfg *Client) PriceChange(id uuid.UUID, record []float64, itemname string, ctx context.Context) {

	if len(record) < 7 {
		cfg.DB.UpdateWeeklyChange(ctx, database.UpdateWeeklyChangeParams{
			ID:         id,
			Weekchange: 0.00,
		})
		return
	}

	weeklyChage := utils.WeeklyPriceChange(record)
	cfg.DB.UpdateWeeklyChange(ctx, database.UpdateWeeklyChangeParams{
		ID:         id,
		Weekchange: weeklyChage,
	})

}
