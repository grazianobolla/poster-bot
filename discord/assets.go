//controls assets lifespan
package discord

import (
	"fmt"
	"time"
)

const MAX_ASSET_LIFE_HOURS = 24
const ASSET_TICKER_FRECUENCY_HOURS = 12

type Asset struct {
	MessageID    string
	ChannelID    string
	AuthorName   string
	CreationTime time.Time
	Uploaded     bool
	Url          string
	Text         string
	Removal      bool
}

//slice containing all current active assets
var assets []Asset

//starts the asset checker ticker
func destroy_ticker() {
	ticker := time.NewTicker(time.Hour * ASSET_TICKER_FRECUENCY_HOURS)
	go check_assets(ticker)
}

//checks if an asset should be removed
func check_assets(ticker *time.Ticker) {
	for {
		<-ticker.C

		var t []Asset //temporal asset slice

		for i := range assets {
			asset := assets[i]
			dif := time.Since(asset.CreationTime)

			if dif < time.Hour*MAX_ASSET_LIFE_HOURS { //if the asset should still exist, its pushed on the temporal slice
				t = append(t, asset)
			} else {
				//adds a alarm clock emoji indicating message timeout
				add_reaction(asset.ChannelID, asset.MessageID, ALARM_CLOCK_EMOJI)
				fmt.Println("Removing asset", asset.Url)
			}
		}

		assets = t
	}
}

//returns true and the asset pointer if a message is an asset
func is_asset(message_id string) (bool, *Asset) {
	for i := range assets {
		asset := &assets[i]

		if asset.MessageID == message_id {
			return true, asset
		}
	}

	return false, nil
}
