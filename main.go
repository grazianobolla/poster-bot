package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"shitposter-bot/database"
	"shitposter-bot/discord"
	"shitposter-bot/instagram"
	"shitposter-bot/shared"
	"shitposter-bot/tenor"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	err := godotenv.Load(os.Args[1])
	if shared.CheckError(err) {
		log.Fatal("Couldn't load env file")
	}
}

func main() {
	discordToken := os.Getenv("DISCORD_TOKEN")

	//social networks TODO: remove this and make it modular
	tenorToken := os.Getenv("TENOR_TOKEN")
	instagramAppId := os.Getenv("IG_APPID")
	instagramAppSecret := os.Getenv("IG_APPSECRET")
	instagramToken := os.Getenv("IG_TOKEN")
	instagramUserId := os.Getenv("IG_USERID")

	database.Start(os.Args[2])
	tenor.Start(tenorToken)
	//go twitter.Start(tw_access_token, tw_access_token_secret, tw_consumer_key, tw_consumer_key_secret)
	go instagram.Start(instagramAppId, instagramAppSecret, instagramToken, instagramUserId)
	go discord.Start(discordToken)

	//wait until we want to stop the program
	chnl := make(chan os.Signal, 1)
	signal.Notify(chnl, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-chnl

	//stop and close
	discord.Stop()
	//twitter.Stop()
	instagram.Stop()
	database.Close()
}
