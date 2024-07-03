package steamapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rpstvs/steamscraper-go/internals/utils"
)

func FetchPlayerData(steamid string) *utils.SteamProfile {

	steamAPIkey := os.Getenv("STEAM_API_KEY")

	Url := fmt.Sprintf(
		"http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=%s&steamids=%s",
		steamAPIkey, steamid,
	)

	resp, _ := http.Get(Url)
	dat, _ := io.ReadAll(resp.Body)
	profile := &utils.SteamProfile{}

	err := json.Unmarshal(dat, profile)

	if err != nil {
		println("coudlnt unmarshal info")
	}

	return profile
}
