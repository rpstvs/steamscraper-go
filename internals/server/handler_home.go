package server

import (
	"net/http"
)

func (srv *Server) home(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("SkinsApp")

	if err == nil && cookie.Value != "" {
		http.Redirect(w, r, "/v1/api/profile", http.StatusPermanentRedirect)
	}

	w.Write([]byte("Welcome to the app"))

}
