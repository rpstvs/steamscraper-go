package steamapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func GetInventory(steamid, last_assetid string, inv *[]utils.Inventory) {
	Url := fmt.Sprintf(
		"https://steamcommunity.com/inventory/%s/730/2?start_assetid=%s", steamid, last_assetid,
	)
	var page utils.Inventory
	resp, err := http.Get(Url)

	if err != nil {
		log.Printf("Error with Inventory Request")
	}

	data, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Println("Error reading response")
		return
	}

	err = json.Unmarshal(data, &page)

	if err != nil {
		log.Println("Couldnt unmarshal response ")
		return
	}

	*inv = append(*inv, page)
	fmt.Println(len(page.Descriptions))
	if page.MoreItems == 1 {
		GetInventory(steamid, page.LastAssetid, inv)
	}
	log.Println("Request successful, inventory being returned")

}
