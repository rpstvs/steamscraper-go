package server

import (
	"fmt"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/database"
)

func (srv *Server) ShowProfile(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt.Println("estamos in")

}

func (srv *Server) ShowBags(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt.Println("estamos in")

}
