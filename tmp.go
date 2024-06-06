package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func (cfg *ApiConfig) WriteToDB(results utils.SearchResult) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	for _, result := range results.Results {
		fmt.Println("vou escrever na base de dados")
		id, err := strconv.Atoi(result.AssetDescription.Classid)
		cfg.DB.CreateItem(ctx, database.CreateItemParams{
			ID:       int32(id),
			Itemname: result.HashName,
		})

		if err != nil {
			log.Println(err)
		}
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
