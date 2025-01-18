package routes

import (
	"youtube-transcriber/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/transcript", handlers.GetTranscriptHandler)
	return r
}
