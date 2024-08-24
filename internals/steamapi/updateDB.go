package steamapi

import (
	"context"
	"fmt"
	"log"
	"strconv"
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
		id, err := strconv.Atoi(result.AssetDescription.Classid)
		//fmt.Println(result.AssetDescription.Classid, int64(id))
		if err != nil {
			log.Println("couldnt convert id")
		}
		image := utils.BuildImageURL(result.AssetDescription.IconURL)

		_, err = cfg.DB.GetItemByName(ctx, result.HashName)
		if err != nil {
			cfg.WriteToDB(result.HashName, image, int64(id), ctx)
		}
		fmt.Println(result.SalePriceText, result.HashName)
		cfg.PriceUpdate(int64(id), result.SalePriceText, ctx)
		cfg.PriceChangeDaily(result.HashName, int64(id))
		cfg.WeeklyPriceChange(result.HashName)

	}

	if start < end {
		start += 100
		fmt.Printf("Dormir 15s - Next Index %d / %d \n", start, end)
		time.Sleep(15 * time.Second)
		cfg.UpdateDB(start)
	}

}

func (cfg *Client) WriteToDB(itemName, url string, id int64, ctx context.Context) {

	_, err := cfg.DB.CreateItem(ctx, database.CreateItemParams{
		ID:         uuid.New(),
		Itemname:   itemName,
		Imageurl:   url,
		Daychange:  0.00,
		Weekchange: 0.00,
		Classid:    id,
	})

	if err != nil {
		log.Print(err)
	}

}

func (cfg *Client) PriceUpdate(id int64, price string, ctx context.Context) {

	priceDb := utils.PriceConverter(price)

	date := utils.ConvertDate()

	_, err := cfg.DB.AddPrice(ctx, database.AddPriceParams{
		ItemClassid: id,
		Pricedate:   date,
		Price:       priceDb,
	})

	if err != nil {
		log.Println(err)
	}

}
