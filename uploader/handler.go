package uploader

import (
	"fmt"
	"shitposter-bot/shared"
	"strings"
)

func UploadAsset(author string, text string, url string) bool {
	mime := shared.GetContentType(url)

	// create caption
	caption := ""
	if text != "" {
		caption = fmt.Sprintf("%s: %s", author, text)
	}

	if strings.Contains(mime, "image") {
		return upload_image(author, text, url, caption)
	} else if strings.Contains(mime, "video") {
		return upload_video(author, text, url, caption)
	}

	return false
}
