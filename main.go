package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient http.Client
}

func main() {

	url := "https://steamcommunity.com/market/search/render/?query=&start=0&norender=1&count=2&search_descriptions=0&sort_column=popular&sort_dir=desc&appid=730"

	steamClient := &Client{
		httpClient: http.Client{},
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("error occurred: %s", err)
		return
	}

	resp, err := steamClient.httpClient.Do(req)

	if err != nil {
		fmt.Printf("error occyrred: %s", err)
		return
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	searchResult := &SearchResult{}

	err = json.Unmarshal(dat, &searchResult)

	if err != nil {
		fmt.Printf("error occurred: %s ", err)
		return
	}

	fmt.Println(searchResult)

}
