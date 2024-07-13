package server

import (
	"net/http"
	"time"

	"github.com/rpstvs/steamscraper-go/internals/auth"
)

func CreateCookie(w http.ResponseWriter, id string) {
	token := auth.CreateToken(id)

	cookie := &http.Cookie{
		Name:     "SkinsApp",
		Value:    token,
		Expires:  time.Now().UTC().Add(24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

}
