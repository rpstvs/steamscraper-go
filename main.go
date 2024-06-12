package main

import (
	"database/sql"
	"fmt"
	"log"
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

	c := cron.New()

	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("starting job")
		cfg.updateDB(0)
	})
	c.Start()
	select {}

}
