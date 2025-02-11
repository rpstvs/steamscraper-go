package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rpstvs/steamscraper-go/internals/server"
)

func main() {

	godotenv.Load(".env")

	server := server.ReturnServer()

	server.Start()

}
