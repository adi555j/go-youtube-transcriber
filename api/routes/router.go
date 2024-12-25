package routes

import (
	"log"
	"youtube-transcriber/api/handlers"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://9b368c3864937ae1f83b5b25a73e9cb9@o4505268028243968.ingest.us.sentry.io/4508529446027264",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for tracing.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}

	r := gin.Default()

	// Once it's done, you can attach the handler as one of your middleware
	r.Use(sentrygin.New(sentrygin.Options{}))

	r.GET("/transcript", handlers.GetTranscriptHandler)
	return r
}
