package main

import (
	"log"
	"youtube-transcriber/api/routes"
	"youtube-transcriber/pkg/config"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Setup and run the router
	r := routes.SetupRouter()

	if err := r.Run(":5002"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
