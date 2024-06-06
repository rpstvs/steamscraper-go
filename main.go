package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rpstvs/steamscraper-go/internals/database"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

type ApiConfig struct {
	steamApiClient steamapi.Client
	DB             *database.Queries
}

const (
	host     = "aws-0-us-east-1.pooler.supabase.com"
	port     = 6543 // default PostgreSQL port
	user     = "postgres.ufbvnnkcvcwqcotbxkcl"
	password = "T8iebtodPi+-314"
	dbname   = "postgres"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	godotenv.Load(".env")
	//dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Printf("no connection to the DB")
	}

	steamClient := steamapi.NewClient(5 * time.Second)
	dbQueries := database.New(db)
	cfg := &ApiConfig{
		steamApiClient: steamClient,
		DB:             dbQueries,
	}

	resultados := cfg.steamApiClient.GetSkins(0)

	cfg.WriteToDB(resultados)

}
