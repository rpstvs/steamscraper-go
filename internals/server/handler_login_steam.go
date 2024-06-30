package server

import (
	"fmt"
	"net/http"

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
	code := r.URL.Query().Get("access_token")
	token, _ := steamOAuthConfig.Exchange(r.Context(), code)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse Steam callback: "+err.Error(), http.StatusInternalServerError)
		return
	}
	steamID := r.Form.Get("openid.identity")
	fmt.Println("Estamos a tratar do login")

	fmt.Println(steamID)

	fmt.Fprintf(w, "Token: %+v", token)
}
