package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rpstvs/steamscraper-go/internals/auth"
)

func (srv *Server) middlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("rpstvs")

		if err != nil {
			println("couldnt get cookie")
			return
		}
		err = auth.ValidateToken(cookie.Value)
		if err != nil {
			fmt.Println("error with token")
			return
		}

		timeNow := time.Now().Unix()

		if cookie.Expires.Unix() <= timeNow {
			fmt.Println("cookie done")
			return
		}

		next.ServeHTTP(w, r)

	}
}
