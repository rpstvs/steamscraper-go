package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML(".market_listing_row_link", func(e *colly.HTMLElement) {

		price := e.ChildText("div.market_listing_price_listings_block > div.market_listing_right_cell.market_listing_their_price > span.market_table_value.normal_price > span.normal_price")
		name := e.ChildText("div.market_listing_item_name_block > .market_listing_item_name")
		fmt.Println(price, name)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".pagebtn", func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.Visit("https://steamcommunity.com/market/search?appid=730#p1_price_asc")
}
