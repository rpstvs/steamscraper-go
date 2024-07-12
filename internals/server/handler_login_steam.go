package server

import (
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

type User struct {
	Name string `json:"name"`
}

func (cfg *Server) loginSteam(w http.ResponseWriter, r *http.Request) {
	fullURL := "http://" + r.Host + r.RequestURI
	urlParsed, _ := r.URL.Parse(fullURL)
	queries, _ := url.ParseQuery(urlParsed.RawQuery)
	openid := queries["openid.claimed_id"][0]
	id := utils.ExtractSteamid(openid)
	profile := steamapi.FetchPlayerData(id)

	name := profile.Response.Players[0].Personaname

	_, err := cfg.DB.GetUserbyId(r.Context(), id)

	if err != nil {
		cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			ID:      uuid.New(),
			Name:    name,
			Steamid: id,
		})

		CreateCookie(w, id)
		http.Redirect(w, r, "/v1/api/profile", http.StatusTemporaryRedirect)
	}

	CreateCookie(w, id)
	http.Redirect(w, r, "/v1/api/profile", http.StatusTemporaryRedirect)

}
