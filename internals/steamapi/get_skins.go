package steamapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func (cfg *Client) GetSkins(start int) utils.SearchResult {
	url := GetUrl(start)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("error occurred: %s", err)
		return utils.SearchResult{}
	}

	resp, err := cfg.httpClient.Do(req)

	if err != nil {
		fmt.Printf("error occyrred: %s", err)
		return utils.SearchResult{}
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error: %s", err)
		return utils.SearchResult{}
	}

	searchResult := &utils.SearchResult{}

	err = json.Unmarshal(dat, &searchResult)

	if err != nil {
		fmt.Printf("error occurred: %s ", err)
		return utils.SearchResult{}
	}
	utils.WriteToFile(*searchResult)
	utils.ParseResults(*searchResult)
	if start < searchResult.TotalCount {
		start += 100
		fmt.Println("sleeping 30 seconds")
		time.Sleep(15 * time.Second)
		fmt.Printf("New Request starting on index: %d \n", start)
		cfg.GetSkins(start)

	}

	return *searchResult
}
