package uploader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"shitposter-bot/database"
	"shitposter-bot/hasher"
	"shitposter-bot/instagram"
	"shitposter-bot/shared"
)

func upload_video(author string, text string, url string, caption string) bool {
	// Encode to base64 and hash, check if already uploaded
	mediaEncodedBase64, base64Ok := retrieve_media_as_base64(url)

	if !base64Ok {
		fmt.Println("failed to retrieve media as base64")
		return false
	}

	// hash and check if already uploaded
	hash := hasher.String2Sha256(mediaEncodedBase64)

	if database.AssetAlreadyUploaded(hash, url) {
		return false
	}

	// post to ig
	instagramPostID, instagramOk := instagram.PostVideo(author, text, url, caption)

	if instagramOk {
		register_upload(author, instagramPostID, hash, url)
		return true
	}

	return false
}

func upload_image(author string, text string, url string, caption string) bool {
	// Encode to base64 and hash, check if already uploaded
	mediaEncodedBase64, base64Ok := retrieve_media_as_base64(url)

	if !base64Ok {
		fmt.Println("failed to retrieve media as base64")
		return false
	}

	// hash and check if already uploaded
	hash := hasher.String2Sha256(mediaEncodedBase64)

	if database.AssetAlreadyUploaded(hash, url) {
		return false
	}

	// post to ig and other media
	instagramPostID, instagramOk := instagram.PostImage(author, text, url, caption)

	if instagramOk {
		register_upload(author, instagramPostID, hash, url)
		return true
	}

	return false
}

func register_upload(author string, post_id string, hash string, url string) {
	media_info := database.MediaInfo{
		Author:    author,
		PostID:    post_id,
		MediaHash: hash,
		MediaURL:  url,
	}

	database.SaveMediaInfo(media_info)
}

// gets a URL image from the web and returns the base64 equivalent
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
