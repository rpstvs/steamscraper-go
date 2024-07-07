package server

import (
	"fmt"
	"net/http"
)

func (srv *Server) showprofile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("estamos in")
	RespondWithJson(w, http.StatusOK, User{
		Name: "Bem jogado",
	})
}
