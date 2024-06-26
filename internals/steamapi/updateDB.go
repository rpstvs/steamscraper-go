package steamapi

import (
	"context"
	"fmt"
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
		fmt.Printf("vou dar update ao item %s \n", result.HashName)
		cfg.WriteToDB(result.HashName, ctx)
		cfg.PriceUpdate(result.HashName, result.SellPrice, ctx)

	}

	if start < end {
		start += 100
		fmt.Println("dormir 15s")
		time.Sleep(15 * time.Second)
		cfg.UpdateDB(start)
	}

}

func (cfg *Client) WriteToDB(itemName string, ctx context.Context) {

	cfg.DB.CreateItem(ctx, database.CreateItemParams{
		ID:       uuid.New(),
		Itemname: itemName,
	})

}

func (cfg *Client) PriceUpdate(itemName string, price int, ctx context.Context) {

	id, err := cfg.DB.GetItemByName(ctx, itemName)
	priceDb := utils.PriceConverter(price)
	if err != nil {
		fmt.Println(err)
	}

	date := utils.ConvertDate()

	cfg.DB.AddPrice(ctx, database.AddPriceParams{
		ItemID:    id,
		Pricedate: date,
		Price:     priceDb,
	})

}
