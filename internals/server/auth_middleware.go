package server

import (
	"fmt"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/auth"
)

func (srv *Server) middlewareAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("rpstvs")

		if err != nil {
			println(err)
			return
		}
		err = auth.ValidateToken(cookie.Value)
		if err != nil {
			fmt.Println("error with token")
			return
		}
		next.ServeHTTP(w, r)

	}
}
