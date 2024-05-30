package steamapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (cfg *Client) GetSkins(start int) SearchResult {
	url := GetUrl(start)
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
	if start < searchResult.TotalCount {
		start += 100
		fmt.Println("sleeping 30 seconds")
		time.Sleep(30 * time.Second)
		fmt.Printf("New Request starting on index: %d \n", start)
		cfg.GetSkins(start)

	}

	return *searchResult
}
