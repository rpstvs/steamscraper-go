package server

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/rpstvs/steamscraper-go/internals/database"
)

type Server struct {
	mux http.ServeMux
	DB  *database.Queries
}

func ReturnServer() *http.Server {
	Port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")
	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	NewServer := Server{
		mux: *http.NewServeMux(),
		DB:  dbQueries,
	}
	cors := middlewareCors(&NewServer.mux)

	NewServer.RegisterEndpoints()

	//NewServer.mux.HandleFunc("GET /v1/api/search", NewServer.GetPrice)

	return &http.Server{
		Addr:    ":" + Port,
		Handler: cors,
	}
}

func (srv *Server) RegisterEndpoints() {

	srv.mux.HandleFunc("/v1/api/price", srv.GetLatestPrice)
	srv.mux.HandleFunc("/v1/api/bag/create", srv.CreateBag)
	srv.mux.HandleFunc("/v1/api/bag/additem", srv.AddItemBag)
	srv.mux.HandleFunc("/v1/api/bag/removeitem", srv.RemoveItemBag)
	srv.mux.HandleFunc("/v1/api/auth/steam", srv.loginSteam)
	srv.mux.HandleFunc("/v1/api/profile", srv.middlewareAuth(srv.showprofile))
}
