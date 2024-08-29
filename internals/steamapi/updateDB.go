package steamapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func (cfg *Client) UpdateDB(index int) {

	resultados := cfg.GetSkins(index)

	start := resultados.Start
	end := resultados.TotalCount

	ctx := context.Background()
	for _, result := range resultados.Results {

		image := utils.BuildImageURL(result.AssetDescription.IconURL)

		id, err := cfg.DB.GetItemByName(ctx, result.HashName)
		if err != nil {
			cfg.WriteToDB(result.HashName, image, ctx)
		}

		cfg.PriceUpdate(id, result.HashName, result.SalePriceText, ctx)
		cfg.PriceChangeDaily(id, result.HashName)
		cfg.WeeklyPriceChange(id, result.HashName)

	}

	if start < end {
		start += 100
		fmt.Printf("Dormir 15s - Next Index %d /%d \n", start, end)
		time.Sleep(15 * time.Second)
		cfg.UpdateDB(start)
	}

}

func (cfg *Client) WriteToDB(itemName, url string, ctx context.Context) {

	_, err := cfg.DB.CreateItem(ctx, database.CreateItemParams{
		ID:         uuid.New(),
		Itemname:   itemName,
		Imageurl:   url,
		Daychange:  0.00,
		Weekchange: 0.00,
	})

	if err != nil {
		log.Print(err)
	}

}

func (cfg *Client) PriceUpdate(id uuid.UUID, itemName string, price string, ctx context.Context) {

	priceDb := utils.PriceConverter(price)

	date := utils.ConvertDate()

	cfg.DB.AddPrice(ctx, database.AddPriceParams{
		ItemID:    id,
		Pricedate: date,
		Price:     priceDb,
	})

}
