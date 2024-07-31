package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
	"github.com/rpstvs/steamscraper-go/internals/server"
	"github.com/rpstvs/steamscraper-go/internals/steamapi"
)

func main() {

	godotenv.Load(".env")

	steamClient := steamapi.NewClient(10 * time.Second)

	server := server.ReturnServer()

	c := cron.New()

	c.AddFunc("0 0 * */1 * *", func() {
		fmt.Println("starting job")
		steamClient.UpdateDB(0)
	})
	c.Start()

	server.Start()

}
