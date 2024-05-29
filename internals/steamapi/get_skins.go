package steamapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (cfg *Client) GetSkins() SearchResult {
	url := GetUrl()
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("error occurred: %s", err)
		return SearchResult{}
	}

	resp, err := cfg.httpClient.Do(req)

	if err != nil {
		fmt.Printf("error occyrred: %s", err)
		return SearchResult{}
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error: %s", err)
		return SearchResult{}
	}

	searchResult := &SearchResult{}

	err = json.Unmarshal(dat, &searchResult)

	if err != nil {
		fmt.Printf("error occurred: %s ", err)
		return SearchResult{}
	}

	for _, item := range searchResult.Results {

		fmt.Println(item.HashName, item.SellPriceText)
	}
	//fmt.Println(searchResult)

	return *searchResult
}
