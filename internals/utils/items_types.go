package utils

type Sticker struct {
	Name       string
	Tournament string
	Condition  string
	Price      float64
}

type Skin struct {
	GunName   string
	SkinName  string
	Condition string
	Price     float64
}

type Case struct {
	CaseName string
	Price    float64
}

type Agent struct {
	Name  string
	Group string
	Price float64
}

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

type SteamProfile struct {
	Response struct {
		Players []struct {
			Steamid                  string `json:"steamid"`
			Communityvisibilitystate int    `json:"communityvisibilitystate"`
			Profilestate             int    `json:"profilestate"`
			Personaname              string `json:"personaname"`
			Profileurl               string `json:"profileurl"`
			Avatar                   string `json:"avatar"`
			Avatarmedium             string `json:"avatarmedium"`
			Avatarfull               string `json:"avatarfull"`
			Avatarhash               string `json:"avatarhash"`
			Lastlogoff               int    `json:"lastlogoff"`
			Personastate             int    `json:"personastate"`
			Primaryclanid            string `json:"primaryclanid"`
			Timecreated              int    `json:"timecreated"`
			Personastateflags        int    `json:"personastateflags"`
			Loccountrycode           string `json:"loccountrycode"`
		} `json:"players"`
	} `json:"response"`
}

type Inventory struct {
	Assets []struct {
		Appid      int    `json:"appid"`
		Contextid  string `json:"contextid"`
		Assetid    string `json:"assetid"`
		Classid    string `json:"classid"`
		Instanceid string `json:"instanceid"`
		Amount     string `json:"amount"`
	} `json:"assets"`
	Descriptions []struct {
		Appid           int    `json:"appid"`
		Classid         string `json:"classid"`
		Instanceid      string `json:"instanceid"`
		Currency        int    `json:"currency"`
		BackgroundColor string `json:"background_color"`
		IconURL         string `json:"icon_url"`
		Descriptions    []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
			Name  string `json:"name"`
			Color string `json:"color,omitempty"`
		} `json:"descriptions"`
		Tradable int `json:"tradable"`
		Actions  []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"actions,omitempty"`
		OwnerDescriptions []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
			Color string `json:"color,omitempty"`
		} `json:"owner_descriptions,omitempty"`
		Name           string `json:"name"`
		NameColor      string `json:"name_color"`
		Type           string `json:"type"`
		MarketName     string `json:"market_name"`
		MarketHashName string `json:"market_hash_name"`
		MarketActions  []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"market_actions,omitempty"`
		Commodity                   int `json:"commodity"`
		MarketTradableRestriction   int `json:"market_tradable_restriction"`
		MarketMarketableRestriction int `json:"market_marketable_restriction"`
		Marketable                  int `json:"marketable"`
		Tags                        []struct {
			Category              string `json:"category"`
			InternalName          string `json:"internal_name"`
			LocalizedCategoryName string `json:"localized_category_name"`
			LocalizedTagName      string `json:"localized_tag_name"`
			Color                 string `json:"color,omitempty"`
		} `json:"tags"`
	} `json:"descriptions"`
	MoreItems           int    `json:"more_items"`
	LastAssetid         string `json:"last_assetid"`
	TotalInventoryCount int    `json:"total_inventory_count"`
	Success             int    `json:"success"`
	Rwgrsn              int    `json:"rwgrsn"`
}

type itemInv struct {
	classid string
	price   float64
}
