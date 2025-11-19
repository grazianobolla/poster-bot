package tenor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"shitposter-bot/shared"
)

// Structs for parsing Tenor API response
type ApiResponse struct {
	Results []Post `json:"results"`
}

type Post struct {
	ID                 string                 `json:"id"`
	ContentDescription string                 `json:"content_description"`
	MediaFormats       map[string]MediaFormat `json:"media_formats"`
}

type MediaFormat struct {
	URL string `json:"url"`
}

var token string
var reg = regexp.MustCompile(`(\d{2,})`)

func Start(t string) {
	token = t
	fmt.Println("Created Tenor")
}

func GetGIFbyURL(url string) (string, bool) {
	fmt.Println("Getting tenor GIF for", url)
	gifId := reg.FindString(url)
	if len(gifId) > 0 {
		url, err := getGifUrl(gifId, token)
		if shared.CheckError(err) {
			return "", false
		}

		return url, true
	}

	return "", false
}

func getGifUrl(id string, token string) (string, error) {
	// Replace with your own (safe) API key via env var if needed
	url := fmt.Sprintf("https://tenor.googleapis.com/v2/posts?ids=%s&key=%s", id, token)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("HTTPRequest error:", err)
	}
	defer resp.Body.Close()

	// Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error:", err)
	}

	// Parse JSON
	var data ApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("JSON parse error:", err)
	}

	// Handle no results
	if len(data.Results) == 0 {
		return "", fmt.Errorf("no results returned.")
	}

	// Extract fields
	post := data.Results[0]

	fmt.Println("ID:", post.ID)
	fmt.Println("Description:", post.ContentDescription)

	// Access GIF (or other formats like tinygif, mp4, mediumgif)
	if gif, ok := post.MediaFormats["tinymp4"]; ok {
		return gif.URL, nil
	}

	return "", fmt.Errorf("error extracting url")
}
