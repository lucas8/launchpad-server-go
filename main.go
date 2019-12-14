package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lucasstettner/launchpad-server/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	} else {
		log.Println("Fetching env variables")
	}

	app.Start()
}
