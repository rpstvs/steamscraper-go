package server

import (
	"fmt"
	"net/http"
)

func (srv *Server) ShowProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("estamos in")

}

func (srv *Server) ShowBags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("estamos in")

}
