package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"shitposter-bot/database"
	"shitposter-bot/discord"
	"shitposter-bot/shared"
	"shitposter-bot/tenor"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	err := godotenv.Load()
	if shared.CheckError(err) {
		log.Fatal("Couldn't load env file")
	}
}

func main() {
	database_path := os.Getenv("DB_PATH")
	discord_token := os.Getenv("DISCORD_TOKEN")

	//social networks TODO: remove this and make it modular
	tenor_token := os.Getenv("TENOR_TOKEN")

	database.Start(database_path)
	tenor.Start(tenor_token)
	//go twitter.Start(tw_access_token, tw_access_token_secret, tw_consumer_key, tw_consumer_key_secret)
	go discord.Start(discord_token)

	//wait until we want to stop the program
	chnl := make(chan os.Signal, 1)
	signal.Notify(chnl, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-chnl

	//stop and close
	discord.Stop()
	//twitter.Stop()
	database.Close()
}
