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

/*
func (cfg *ApiConfig) PriceUpdate(results utils.SearchResult) {
	ctx := context.Background()

	for _, result := range results.Results {
		fmt.Println("vou escrever na base de dados")

		cfg.DB.CreateItem(ctx, database.{
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
*/
