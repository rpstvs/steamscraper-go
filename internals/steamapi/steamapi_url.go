package steamapi

func GetUrl() string {

	baseUrl := "https://steamcommunity.com/market/search/render/?query=&start=100&count=30&search_descriptions=0&norender=1&sort_column=popular&sort_dir=desc&appid=730"

	return baseUrl
}
