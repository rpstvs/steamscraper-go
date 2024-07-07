package server

import (
	"net/http"
	"time"

	"github.com/rpstvs/steamscraper-go/internals/auth"
)

func CreateCookie(w http.ResponseWriter, name, id string) {
	token := auth.CreateToken(id)

	cookie := &http.Cookie{

		Name:    name,
		Value:   token,
		Expires: time.Now().UTC().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)

	return
}
