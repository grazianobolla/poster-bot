package instagram

import (
	"fmt"
	"shitposter-bot/shared"
	"time"

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

func PostImage(author string, text string, url string, caption string) (string, bool) {
	containerId, err := createContainerImage(url, caption)
	fmt.Println("Ig: generated containerId for image ", containerId, " waiting 10 seconds...")
	time.Sleep(10 * time.Second)

	if shared.CheckError(err) {
		return "", false
	}

	err = publishMedia(containerId)

	if !shared.CheckError(err) {
		fmt.Println("Ig: Posted image to Instagram", author, text, url, caption)
		return containerId, true
	}

	return "", false
}

func PostVideo(author string, text string, url string, caption string) (string, bool) {
	containerId, err := createContainerVideo(url, caption)
	fmt.Println("Ig: generated containerId for video ", containerId, " waiting 30 seconds...")

	for i := 0; i < 10; i++ {
		time.Sleep(30 * time.Second)

		status, err := checkMediaStatus(containerId)

		if shared.CheckError(err) {
			return "", false
		}

		fmt.Println("Checking status for", containerId, "result:", status)

		if status == "FINISHED" {
			break
		}

		if status == "ERROR" {
			return "", false
		}
	}

	if shared.CheckError(err) {
		return "", false
	}

	err = publishMedia(containerId)

	if !shared.CheckError(err) {
		fmt.Println("Ig: Posted video to Instagram", author, text, url, caption)
		return containerId, true
	}

	return "", false
}
