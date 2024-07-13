package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rs/cors"
)

type Server struct {
	mux  http.ServeMux
	DB   *database.Queries
	Port string
}

func ReturnServer() *Server {
	Port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")
	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	NewServer := Server{

		DB:   dbQueries,
		Port: Port,
	}
	//cors := middlewareCors(&NewServer.mux)

	return &NewServer
}

func (srv *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", srv.home)
	mux.HandleFunc("/v1/api/price", srv.GetLatestPrice)
	mux.HandleFunc("/v1/api/bag/create", srv.middlewareAuth(srv.CreateBag))
	mux.HandleFunc("/v1/api/bag/additem", srv.middlewareAuth(srv.AddItemBag))
	mux.HandleFunc("/v1/api/bag/removeitem", srv.middlewareAuth(srv.RemoveItemBag))
	mux.HandleFunc("/v1/api/auth/steam", srv.loginSteamCallback)
	mux.HandleFunc("/v1/api/profile", srv.middlewareAuth(srv.ShowProfile))
	mux.HandleFunc("/v1/api/bags", srv.middlewareAuth(srv.ShowBags))

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := c.Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func (srv *Server) RegisterEndpoints() {
	srv.mux.HandleFunc(".", srv.home)
	srv.mux.HandleFunc("/v1/api/price", srv.GetLatestPrice)
	srv.mux.HandleFunc("/v1/api/bag/create", srv.middlewareAuth(srv.CreateBag))
	srv.mux.HandleFunc("/v1/api/bag/additem", srv.middlewareAuth(srv.AddItemBag))
	srv.mux.HandleFunc("/v1/api/bag/removeitem", srv.middlewareAuth(srv.RemoveItemBag))
	srv.mux.HandleFunc("/v1/api/auth/steam", srv.loginSteamCallback)
	srv.mux.HandleFunc("/v1/api/profile", srv.middlewareAuth(srv.ShowProfile))
	srv.mux.HandleFunc("/v1/api/bags", srv.middlewareAuth(srv.ShowBags))
}
