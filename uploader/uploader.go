package uploader

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"shitposter-bot/database"
	"shitposter-bot/hasher"
	"shitposter-bot/shared"
	"shitposter-bot/twitter"
)

func UploadAsset(author string, text string, url string) bool {
	if media_encoded_base64, ok := retrieve_media_as_base64(url); ok {
		//check hash
		hash := hasher.String2Sha256(media_encoded_base64)
		if database.AssetAlreadyUploaded(hash, url) {

			return false
		}

		if tweet_id, ok := twitter.TweetImage(author, text, media_encoded_base64); ok {
			//the image hasnt been uploaded, we save it to the database
			media_info := database.MediaInfo{
				Author:    author,
				TweetID:   tweet_id,
				MediaHash: hash,
				MediaURL:  url,
			}

			database.SaveMediaInfo(media_info)

			fmt.Println("Tweeted image", url, text)
			return true
		}
	}

	return false
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

	return to_base64(bytes), true
}

//byte[] to base64 string
func to_base64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
