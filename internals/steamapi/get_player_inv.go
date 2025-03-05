package steamapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func GetInventory(steamid string) utils.Inventory {
	Url := fmt.Sprintf(
		"https://steamcommunity.com/inventory/%s/730/2?count=5000", steamid,
	)
	var page utils.Inventory
	resp, err := http.Get(Url)

	if err != nil {
		log.Printf("Error with Inventory Request")
		return utils.Inventory{}
	}

	data, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Println("Error reading response")
		return utils.Inventory{}
	}

	err = json.Unmarshal(data, &page)

	if err != nil {
		log.Println("Couldnt unmarshal response ")
		return utils.Inventory{}
	}

	log.Println("Request successful, inventory being returned")
	return page
}
