package main

import (
	"log"
	"youtube-transcriber/api/routes"
	"youtube-transcriber/pkg/config"

	"github.com/getsentry/sentry-go"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://9b368c3864937ae1f83b5b25a73e9cb9@o4505268028243968.ingest.us.sentry.io/4508529446027264",
		EnableTracing:    true,
		TracesSampleRate: 0.3,
	}); err != nil {
		log.Fatalf("Sentry initialization failed: %v", err)
	}
	// Ensure Sentry events are flushed before the application shuts down
	defer sentry.Flush(2)

	// Setup and run the router
	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
