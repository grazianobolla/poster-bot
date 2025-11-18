package instagram

import (
	"errors"
	"fmt"
	"shitposter-bot/shared"

	fb "github.com/huandu/facebook/v2"
)

var session *fb.Session
var instagramId string

func publishMedia(mediaId string) error {
	result, err := session.Post(
		fmt.Sprintf("/%s/media_publish", instagramId),
		fb.Params{"creation_id": mediaId},
	)

	if shared.CheckError(err) {
		return err
	}

	ok := result.GetField("id")

	if ok == nil {
		return errors.New("couldn't publish media")
	}

	return nil
}

func createContainerImage(url string) (string, error) {
	result, err := session.Post(
		fmt.Sprintf("/%s/media", instagramId),
		fb.Params{"image_url": url},
	)

	if shared.CheckError(err) {
		return "", err
	}

	mediaId := result.GetField("id").(string)
	return mediaId, nil
}

func createContainerVideo(url string) (string, error) {
	result, err := session.Post(
		fmt.Sprintf("/%s/media", instagramId),
		fb.Params{"video_url": url, "media_type": "REELS"},
	)

	if shared.CheckError(err) {
		return "", err
	}

	mediaId := result.GetField("id").(string)
	return mediaId, nil
}
