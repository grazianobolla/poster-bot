package twitter

import (
	"fmt"
	"net/url"
	"shitposter-bot/shared"

	"github.com/ChimeraCoder/anaconda"
)

var client *anaconda.TwitterApi

func post_tweet(mediaId string, text string) (anaconda.Tweet, bool) {
	//post the tweet
	v := url.Values{}
	v.Set("media_ids", mediaId)
	tweet, e := client.PostTweet(text, v)

	if !shared.CheckError(e) {
		fmt.Println("Posted media to Twitter ", mediaId)
	}

	return tweet, !shared.CheckError(e)
}
