package uploader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"shitposter-bot/database"
	"shitposter-bot/hasher"
	"shitposter-bot/shared"
	"shitposter-bot/twitter"
)

func upload_video(author string, text string, url string) bool {
	if media_bytes, ok := retrieve_data_as_bytes(url); ok {
		//check hash
		hash := hasher.Byte2Sha256(media_bytes)
		if database.AssetAlreadyUploaded(hash, url) {
			return false
		}

		if tweet_id, ok := twitter.TweetVideo(author, text, media_bytes); ok {
			//the video hasnt been uploaded, we save it to the database
			register_upload(author, tweet_id, hash, url)
			fmt.Println("Tweeted video", hash, text)
			return true
		}
	}

	return false
}

func upload_image(author string, text string, url string) bool {
	if media_encoded_base64, ok := retrieve_media_as_base64(url); ok {
		//check hash
		hash := hasher.String2Sha256(media_encoded_base64)
		if database.AssetAlreadyUploaded(hash, url) {
			return false
		}

		if tweet_id, ok := twitter.TweetImage(author, text, media_encoded_base64); ok {
			//the image hasnt been uploaded, we save it to the database
			register_upload(author, tweet_id, hash, url)
			fmt.Println("Tweeted image", hash, text)
			return true
		}
	}

	return false
}

func register_upload(author string, tweet_id int64, hash string, url string) {
	media_info := database.MediaInfo{
		Author:    author,
		TweetID:   tweet_id,
		MediaHash: hash,
		MediaURL:  url,
	}

	database.SaveMediaInfo(media_info)
}

//gets a URL image from the web and returns the base64 equivalent
func retrieve_media_as_base64(url string) (string, bool) {
	resp, err := http.Get(url)
	if shared.CheckError(err) || resp.StatusCode != http.StatusOK {
		return "", false
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if shared.CheckError(err) {
		return "", false
	}

	return shared.ToBase64(bytes), true
}

func retrieve_data_as_bytes(url string) ([]byte, bool) {
	resp, err := http.Get(url)
	if shared.CheckError(err) || resp.StatusCode != http.StatusOK {
		return nil, false
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if shared.CheckError(err) {
		return nil, false
	}

	return bytes, true
}
