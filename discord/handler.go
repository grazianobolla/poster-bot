package discord

import "fmt"

func Start(token string, min_reactions_needed int) {
	minReactionsNeeded = min_reactions_needed
	go destroy_ticker()
	start_connection(token)
}

func Stop() {
	client.Close()
	fmt.Println("Shitposter Bot Discord stopped running")
}
