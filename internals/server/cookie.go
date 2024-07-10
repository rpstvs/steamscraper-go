package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rpstvs/steamscraper-go/internals/auth"
)

func CreateCookie(w http.ResponseWriter, id string) {
	token := auth.CreateToken(id)

	cookie := &http.Cookie{

		Name:    "SkinsApp",
		Value:   token,
		Expires: time.Unix(time.Now().UTC().Add(24*time.Hour).Unix(), 0),
	}
	fmt.Println(cookie)

	http.SetCookie(w, cookie)

}
