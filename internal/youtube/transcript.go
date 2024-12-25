package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Transcript struct {
	Text     string
	Duration float64
	Offset   float64
	Lang     string
}

// FetchTranscript retrieves the full transcript of a YouTube video.
func FetchTranscript(videoID, lang string) (string, string, error) {
	log.Println("🔍 Fetching transcript for videoID:", videoID)

	videoID = extractVideoID(videoID)
	log.Println("✅ Extracted Video ID:", videoID)

	videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
	resp, err := http.Get(videoURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch video page: %w", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(body), `"captionTracks":`) {
		log.Println("❌ No captions found in the video page response.")
		return "", "", fmt.Errorf("no captions available for video %s", videoID)
	}

	// Extract captionTracks JSON
	re := regexp.MustCompile(`"captionTracks":(\[.*?\])`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		log.Println("❌ No captionTracks found in the video response.")
		return "", "", fmt.Errorf("no captionTracks found")
	}

	var captionTracks []map[string]interface{}
	err = json.Unmarshal([]byte(matches[1]), &captionTracks)
	if err != nil {
		log.Println("❌ Failed to parse captionTracks JSON:", err)
		return "", "", fmt.Errorf("failed to parse captionTracks")
	}

	// Find the appropriate transcript URL
	var transcriptURL string
	for _, track := range captionTracks {
		if lang == "auto" || track["languageCode"] == lang {
			if baseURL, ok := track["baseUrl"].(string); ok {
				decodedURL, err := url.QueryUnescape(baseURL)
				if err != nil {
					log.Println("❌ Failed to decode baseUrl:", err)
					return "", "", fmt.Errorf("failed to decode transcript URL: %w", err)
				}
				transcriptURL = decodedURL
				break
			}
		}
	}

	if transcriptURL == "" {
		log.Println("❌ No valid transcript URL found for language:", lang)
		return "", "", fmt.Errorf("no transcript available for language: %s", lang)
	}

	log.Println("✅ Transcript URL:", transcriptURL)

	// Fetch the transcript
	transcriptResp, err := http.Get(transcriptURL)
	if err != nil {
		log.Println("❌ Failed to fetch transcript from URL.")
		return "", "", fmt.Errorf("failed to fetch transcript: %w", err)
	}
	defer transcriptResp.Body.Close()

	transcriptBody, err := ioutil.ReadAll(transcriptResp.Body)
	if err != nil {
		log.Println("❌ Failed to read transcript response body.")
		return "", "", fmt.Errorf("failed to read transcript response: %w", err)
	}

	if len(transcriptBody) == 0 {
		log.Println("❗ Transcript body is empty. Retrying with 'auto' language selection...")
		return FetchTranscript(videoID, "auto")
	}

	// Extract transcript content
	reXML := regexp.MustCompile(`<text start="([^"]*)" dur="([^"]*)">([^<]*)<\/text>`)
	matchesXML := reXML.FindAllStringSubmatch(string(transcriptBody), -1)
	log.Println("🔄 Found Transcript Matches:", len(matchesXML))

	var transcriptText []string
	for _, match := range matchesXML {
		if len(match) < 4 {
			continue
		}
		transcriptText = append(transcriptText, match[3])
	}

	combinedTranscript := strings.Join(transcriptText, " ")
	log.Println("✅ Full Transcript Text Ready")
	return combinedTranscript, "Sample Video Title", nil
}

// Extract YouTube Video ID
func extractVideoID(videoID string) string {
	re := regexp.MustCompile(`(?:v=|youtu\.be/|embed/|watch\?v=)([^"&?/ ]{11})`)
	matches := re.FindStringSubmatch(videoID)
	if len(matches) > 1 {
		return matches[1]
	}
	return videoID
}
