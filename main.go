package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"shitposter-bot/database"
	"shitposter-bot/discord"
	"shitposter-bot/tenor"
	"shitposter-bot/twitter"
	"syscall"
)

func main() {
	var discord_token, database_name string
	var access_token, access_token_secret, consumer_key, consumer_key_secret, tenor_token string

	flag.StringVar(&discord_token, "d", "", "Discord Token")
	flag.StringVar(&discord_token, "db", "", "Database Name")

	//TODO: modularizar las redes sociales
	flag.StringVar(&access_token, "ta", "", "Twitter Access Token")
	flag.StringVar(&access_token_secret, "tas", "", "Twitter Access Token Secret")
	flag.StringVar(&consumer_key, "tc", "", "Twitter Consumer Token")
	flag.StringVar(&consumer_key_secret, "tcs", "", "Twitter Consumer Token Secret")
	flag.StringVar(&tenor_token, "tt", "", "Tenor Token")
	flag.Parse()

	if discord_token == "" || access_token == "" || access_token_secret == "" || consumer_key == "" || consumer_key_secret == "" || tenor_token == "" {
		fmt.Println("Missing tokens")
		return
	}

	database.Start(database_name)
	tenor.Start(tenor_token)
	go twitter.Start(access_token, access_token_secret, consumer_key, consumer_key_secret)
	go discord.Start(discord_token)

	//wait until we want to stop the program
	chnl := make(chan os.Signal, 1)
	signal.Notify(chnl, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-chnl

	//stop and close
	discord.Stop()
	twitter.Stop()
	database.Close()
}
