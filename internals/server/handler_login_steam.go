package server

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
	"github.com/rpstvs/steamscraper-go/internals/utils"
)

type User struct {
	Name string `json:"name"`
}

func (cfg *Server) loginSteamCallback(w http.ResponseWriter, r *http.Request) {
	fullURL := "http://" + r.Host + r.RequestURI

	urlParsed, _ := url.Parse(fullURL)

	params := urlParsed.Query()

	validationParams := url.Values{}
	validationParams.Add("openid.assoc_handle", params.Get("openid.assoc_handle"))
	validationParams.Add("openid.signed", params.Get("openid.signed"))
	validationParams.Add("openid.sig", params.Get("openid.sig"))
	validationParams.Add("openid.ns", params.Get("openid.ns"))
	validationParams.Add("openid.mode", "check_authentication")
	for _, param := range strings.Split(params.Get("openid.signed"), ",") {
		validationParams.Add("openid."+param, params.Get("openid."+param))
	}

	resp, err := http.PostForm("https://steamcommunity.com/openid/login", validationParams)
	if err != nil {
		http.Error(w, "Error validating OpenID response", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	if !strings.Contains(string(body), "is_valid:true") {
		http.Error(w, "OpenID validation failed", http.StatusUnauthorized)
		return
	}

	queries, _ := url.ParseQuery(urlParsed.RawQuery)
	openid := queries["openid.claimed_id"][0]
	id := utils.ExtractSteamid(openid)
	profile := steamapi.FetchPlayerData(id)

	name := profile.Response.Players[0].Personaname

	_, err = cfg.DB.GetUserbyId(r.Context(), id)

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
