package main

import (
	"bms-movies/app"
	"bms-movies/server"
	"log"
)

func main() {

	err := app.Initialize()
	if err != nil {
		log.Fatalf("App could not start %s", err)
	}
	server.InitializeServer()
}
