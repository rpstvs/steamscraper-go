package steamapi

type SearchResult struct {
	Success    bool `json:"success"`
	Start      int  `json:"start"`
	Pagesize   int  `json:"pagesize"`
	TotalCount int  `json:"total_count"`
	Searchdata struct {
		Query              string `json:"query"`
		SearchDescriptions bool   `json:"search_descriptions"`
		TotalCount         int    `json:"total_count"`
		Pagesize           int    `json:"pagesize"`
		Prefix             string `json:"prefix"`
		ClassPrefix        string `json:"class_prefix"`
	} `json:"searchdata"`
	Results []struct {
		Name             string `json:"name"`
		HashName         string `json:"hash_name"`
		SellListings     int    `json:"sell_listings"`
		SellPrice        int    `json:"sell_price"`
		SellPriceText    string `json:"sell_price_text"`
		AppIcon          string `json:"app_icon"`
		AppName          string `json:"app_name"`
		AssetDescription struct {
			Appid           int    `json:"appid"`
			Classid         string `json:"classid"`
			Instanceid      string `json:"instanceid"`
			Currency        int    `json:"currency"`
			BackgroundColor string `json:"background_color"`
			IconURL         string `json:"icon_url"`
			IconURLLarge    string `json:"icon_url_large"`
			Descriptions    []struct {
				Type  string `json:"type"`
				Value string `json:"value"`
				Color string `json:"color,omitempty"`
			} `json:"descriptions"`
			Tradable                    int    `json:"tradable"`
			Name                        string `json:"name"`
			NameColor                   string `json:"name_color"`
			Type                        string `json:"type"`
			MarketName                  string `json:"market_name"`
			MarketHashName              string `json:"market_hash_name"`
			Commodity                   int    `json:"commodity"`
			MarketTradableRestriction   int    `json:"market_tradable_restriction"`
			Marketable                  int    `json:"marketable"`
			MarketBuyCountryRestriction string `json:"market_buy_country_restriction"`
		} `json:"asset_description"`
		SalePriceText string `json:"sale_price_text"`
	} `json:"results"`
}
