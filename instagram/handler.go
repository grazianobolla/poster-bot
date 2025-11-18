package instagram

import (
	"fmt"
	"shitposter-bot/shared"

	fb "github.com/huandu/facebook/v2"
)

// starts the instagram client
func Start(appId, appSecret, accessToken, igID string) {
	session = fb.New(appId, appSecret).Session(accessToken)
	session.BaseURL = "https://graph.instagram.com/v24.0/"
	instagramId = igID
	fmt.Println("Shitposter Bot Instagram now running")
}

// stops the instagram client
func Stop() {
	fmt.Println("Shitposter Bot Twitter stopped running")
}

func PostImage(author string, text string, url string) (string, bool) {
	mediaId, err := createContainerImage(url)

	if shared.CheckError(err) {
		return "", false
	}

	err = publishMedia(mediaId)

	if !shared.CheckError(err) {
		fmt.Println("Posted image to Instagram", author, text, url)
		return mediaId, true
	}

	return "", false
}

func PostVideo(author string, text string, url string) (string, bool) {
	mediaId, err := createContainerVideo(url)

	if shared.CheckError(err) {
		return "", false
	}

	err = publishMedia(mediaId)

	if !shared.CheckError(err) {
		fmt.Println("Posted video to Instagram", author, text, url)
		return mediaId, true
	}

	return "", false
}
