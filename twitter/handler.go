package twitter

import (
	"fmt"
	"shitposter-bot/shared"

	"github.com/ChimeraCoder/anaconda"
)

//starts the twitter client
func Start(access_token string, access_token_secret string, consumer_key string, consumer_key_secret string) {
	client = anaconda.NewTwitterApiWithCredentials(access_token, access_token_secret, consumer_key, consumer_key_secret)
	fmt.Println("Shitposter Bot Twitter is now running")
}

//stops the twitter client
func Stop() {
	client.Close()
	fmt.Println("Shitposter Bot Twitter stopped running")
}

func TweetImage(author string, text string, base64media string) (int64, bool) {
	//upload image to tweeter
	media, err := client.UploadMedia(base64media)
	if shared.CheckError(err) {
		return 0, false
	}

	//post the tweet
	tweet, ok := post_tweet(media.MediaIDString, text)

	if ok {
		return tweet.Id, true
	}

	return 0, false
}

func TweetVideo(author string, text string, data []byte) (int64, bool) {
	size_bytes := len(data)

	media, err := client.UploadVideoInit(size_bytes, "video/mp4")

	if shared.CheckError(err) {
		return 0, false
	}

	chunk_idx := 0

	//separates the video in 500000bytes chunks and uploads them one at the time
	for i := 0; i < size_bytes; i += 500000 {
		fmt.Println("Uploading video chunk", chunk_idx)

		chunk_size := 500000

		if i+chunk_size > size_bytes {
			chunk_size = size_bytes - i
		}

		err = client.UploadVideoAppend(media.MediaIDString, chunk_idx, shared.ToBase64(data[i:i+chunk_size]))

		if shared.CheckError(err) {
			fmt.Println("Error uploading chunk")
			return 0, false
		}

		chunk_idx++
	}

	video, err := client.UploadVideoFinalize(media.MediaIDString)

	if shared.CheckError(err) {
		return 0, false
	}

	//post the tweet
	tweet, ok := post_tweet(video.MediaIDString, text)

	if ok {
		return tweet.Id, true
	}

	return 0, false
}
