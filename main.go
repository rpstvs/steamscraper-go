package main

import (
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
	"github.com/rpstvs/steamscraper-go/internals/server"
)

func main() {

	godotenv.Load(".env")

	//steamClient := steamapi.NewClient(10 * time.Second)

	server := server.ReturnServer()

	c := cron.New()

	c.AddFunc("*1 * * * * *", func() {
		fmt.Println("starting job")

	})
	c.Start()

	//steamClient.UpdateDB(0)
	server.ListenAndServe()

}
