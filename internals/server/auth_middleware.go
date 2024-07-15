package server

import (
	"fmt"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/auth"
	"github.com/rpstvs/steamscraper-go/internals/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (srv *Server) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("SkinsApp")

		if err != nil {
			println("couldnt get cookie")
			return
		}
		err = auth.ValidateToken(cookie.Value)
		if err != nil {
			fmt.Println("error with token")
			return
		}

		steamid := auth.GetSubject(cookie.Value)

		user, _ := srv.DB.GetUserbyId(r.Context(), steamid)

		handler(w, r, user)

	}
}
