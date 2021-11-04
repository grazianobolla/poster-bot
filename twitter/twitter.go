package twitter

import (
	"fmt"
	"net/url"
	"shitposter-bot/shared"

	"github.com/ChimeraCoder/anaconda"
)

var client *anaconda.TwitterApi

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

//tweets asset, coming from the URL: has to be a file
func TweetImage(author string, text string, base64media string) (int64, bool) {
	//upload image to tweeter
	media, err := client.UploadMedia(base64media)
	if shared.CheckError(err) {
		return 0, false
	}

	v := url.Values{}
	v.Set("media_ids", media.MediaIDString)

	//post the tweet
	tweet, e := client.PostTweet(" ", v)
	if shared.CheckError(e) {
		return 0, false
	}

	return tweet.Id, true
}
