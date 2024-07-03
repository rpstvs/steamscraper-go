package server

import (
	"net/http"
	"net/url"

	"github.com/rpstvs/steamscraper-go/internals/steamapi"
	"golang.org/x/oauth2"
)

var (
	steamOAuthConfig = oauth2.Config{
		ClientID:    "__-",
		RedirectURL: "http://localhost:8080/auth/steam/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://steamcommunity.com/openid/login",
			TokenURL: "https://steamcommunity.com/openid/login",
		},
		Scopes: []string{},
	}
)

func (cfg *Server) loginSteam(w http.ResponseWriter, r *http.Request) {
	fullURL := "http://" + r.Host + r.RequestURI
	urlParsed, _ := r.URL.Parse(fullURL)
	queries, _ := url.ParseQuery(urlParsed.RawQuery)
	id := queries["openid.claimed_id"][0]

	profile := steamapi.FetchPlayerData(id)

	name := profile.Response.Players[0].Personaname
}

//http://localhost:8080/v1/api/auth/steam?openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.mode=id_res&openid.op_endpoint=https%3A%2F%2Fsteamcommunity.com%2Fopenid%2Flogin&openid.claimed_id=https%3A%2F%2Fsteamcommunity.com%2Fopenid%2Fid%2F76561198000318676&openid.identity=https%3A%2F%2Fsteamcommunity.com%2Fopenid%2Fid%2F76561198000318676&openid.return_to=http%3A%2F%2Flocalhost%3A8080%2Fv1%2Fapi%2Fauth%2Fsteam&openid.response_nonce=2024-06-30T18%3A11%3A11ZqRt6FJABqlB%2BVowcPds03mcaDjk%3D&openid.assoc_handle=1234567890&openid.signed=signed%2Cop_endpoint%2Cclaimed_id%2Cidentity%2Creturn_to%2Cresponse_nonce%2Cassoc_handle&openid.sig=54%2FH3ptnfJutyx8LMXRHcE3qvwI%3D

//Utilizador faz login
// retirar o nome da steam
// registar o user na DB
