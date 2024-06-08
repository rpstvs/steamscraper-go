package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func (cfg *ApiConfig) WriteToDB(results utils.SearchResult) {

	ctx := context.Background()

	for _, result := range results.Results {
		//fmt.Println("vou escrever na base de dados")

		cfg.DB.CreateItem(ctx, database.CreateItemParams{
			ID:       uuid.New(),
			Itemname: result.HashName,
		})

	}

	start := results.Start
	if start < results.TotalCount {
		start += 100
		fmt.Println("sleeping 30 seconds")
		time.Sleep(15 * time.Second)
		fmt.Printf("New Request starting on index: %d \n", start)
		resultados := cfg.steamApiClient.GetSkins(start)
		cfg.WriteToDB(resultados)
	}
}

func (cfg *ApiConfig) PriceUpdate(results utils.SearchResult) {
	ctx := context.Background()

	for _, result := range results.Results {
		fmt.Printf("Vou adicionar o preço de %s\n", result.HashName)
		id, err := cfg.DB.GetItemByName(ctx, result.HashName)
		price := utils.PriceConverter(result.SellPrice)
		if err != nil {
			fmt.Println(err)
		}

		cfg.DB.AddPrice(ctx, database.AddPriceParams{
			ItemID:    id,
			Pricedate: time.Now().UTC(),
			Price:     price,
		})

	}

	start := results.Start
	if start < results.TotalCount {
		start += 100
		fmt.Println("sleeping 30 seconds")
		time.Sleep(15 * time.Second)
		fmt.Printf("New Request starting on index: %d \n", start)
		resultados := cfg.steamApiClient.GetSkins(start)
		cfg.WriteToDB(resultados)
	}
}
