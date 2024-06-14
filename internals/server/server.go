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

func ReturnServer(port string) *http.Server {

	dbURL := os.Getenv("DATABASE_URL")
	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	NewServer := Server{
		mux: *http.NewServeMux(),
		DB:  dbQueries,
	}

	NewServer.mux.HandleFunc("GET /v1/api/search", NewServer.GetPrice)

	return &http.Server{
		Addr:    ":" + port,
		Handler: middlewareCors(&NewServer.mux),
	}
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, POST, DELETE ")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "Options" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
