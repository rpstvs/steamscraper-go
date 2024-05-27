package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML(".market_listing_row_link", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		//name := e.ChildText(".market_listing_row_link")
		//fmt.Println(name)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://steamcommunity.com/market/search?appid=730#p1_price_asc")
}
