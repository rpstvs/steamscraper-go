package main

import (
	"fmt"
	"time"

	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

type config struct {
	steamApiClient steamapi.Client
}

func main() {

	steamClient := steamapi.NewClient(5 * time.Second)

	cfg := &config{
		steamApiClient: steamClient,
	}

	Resultados := cfg.steamApiClient.GetSkins(7500)

	for _, resultado := range Resultados.Results {
		fmt.Println(resultado.HashName, resultado.SellPriceText)
	}

}
