package handlers

import (
	"log"
	"net/http"
	"youtube-transcriber/internal/youtube"

	"github.com/gin-gonic/gin"
)

// @Summary Get YouTube Transcript
// @Description Fetch full transcript text from a YouTube video
// @Param videoId query string true "YouTube Video ID"
// @Param lang query string false "Language Code (e.g., en, auto)"
// @Produce plain
// @Success 200 {string} string "Full Transcript Text"
// @Router /transcript [get]

func GetTranscriptHandler(c *gin.Context) {
	videoID := c.Query("videoId")
	lang := c.DefaultQuery("lang", "en")

	if videoID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "videoId is required"})
		return
	}

	log.Println("📥 Received request for video:", videoID)

	transcript, _, err := youtube.FetchTranscript(videoID, lang)
	if err != nil {
		log.Println("❌ Error fetching transcript:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if transcript == "" {
		log.Println("❗ No transcript found for video:", videoID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Transcript not found"})
		return
	}

	log.Println("✅ Full Transcript successfully fetched for:", videoID)
	c.String(http.StatusOK, transcript)
}
