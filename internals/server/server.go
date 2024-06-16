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

	NewServer.RegisterEndpoints()

	//NewServer.mux.HandleFunc("GET /v1/api/search", NewServer.GetPrice)

	return &http.Server{
		Addr:    ":" + Port,
		Handler: middlewareCors(&NewServer.mux),
	}
}

func (srv *Server) RegisterEndpoints() {

	srv.mux.HandleFunc("GET /v1/api/price/latest", srv.GetLatestPrice)
}
