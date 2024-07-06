package server

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/auth"
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
	fmt.Println(name, id)

	cfg.DB.GetUserbyId(r.Context(), id)

	cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:      uuid.New(),
		Name:    name,
		Steamid: id,
	})

	token := auth.CreateToken(id)

	cookie := &http.Cookie{
		Name:    name,
		Value:   token,
		Expires: time.Now().UTC().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)

	fmt.Println(cookie.Expires)

	RespondWithJson(w, http.StatusOK, User{
		Name: name,
	})

}
