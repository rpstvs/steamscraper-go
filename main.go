package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/rpstvs/steamscraper-go/internals/server"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

func main() {

	godotenv.Load(".env")
	Port := os.Getenv("PORT")
	steamClient := steamapi.NewClient(10 * time.Second)

	server := server.ReturnServer(Port)
	c := cron.New()

	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("starting job")
		steamClient.UpdateDB(0)
	})
	c.Start()
	server.ListenAndServe()

}
