package main

import (
	"log"
	"user-ratings/initialize"
)

func main() {
	err := initialize.Init()
	if err != nil {
		log.Fatalf("App could not start %s", err)
	}
}
