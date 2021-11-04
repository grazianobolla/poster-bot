package discord

import "fmt"

func Start(token string) {
	go destroy_ticker()
	start_connection(token)
}

func Stop() {
	client.Close()
	fmt.Println("Shitposter Bot Discord stopped running")
}
