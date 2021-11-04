package uploader

import (
	"shitposter-bot/shared"
	"strings"
)

func UploadAsset(author string, text string, url string) bool {
	mime := shared.GetContentType(url)

	if strings.Contains(mime, "image") {
		return upload_image(author, text, url)
	} else if strings.Contains(mime, "video") {
		return upload_video(author, text, url)
	}

	return false
}
