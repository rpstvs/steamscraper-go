package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

type ApiConfig struct {
	steamApiClient steamapi.Client
	DB             *database.Queries
}

func main() {

	godotenv.Load(".env")
	Port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Printf("no connection to the DB")
	}
	steamClient := steamapi.NewClient(10 * time.Second)
	dbQueries := database.New(db)
	cfg := &ApiConfig{
		steamApiClient: steamClient,
		DB:             dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/search", cfg.GetPrice)
	cors := middlewareCors(mux)

	server := &http.Server{
		Addr:    ":" + Port,
		Handler: cors,
	}

	c := cron.New()

	c.AddFunc("* * */1 * *", func() {
		fmt.Println("starting job")
		cfg.updateDB(0)
	})
	c.Start()
	server.ListenAndServe()

}
