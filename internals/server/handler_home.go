package server

import (
	"fmt"
	"net/http"
)

func (srv *Server) home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("helloworld")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))

}
