package database

import (
	"database/sql"
	"fmt"
	"log"
	"shitposter-bot/shared"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

type MediaInfo struct {
	Author    string
	TweetID   int64
	MediaHash string
	MediaURL  string
}

//starts the database file connection
func Start(database_path string) {
	d, err := sql.Open("sqlite3", database_path)
	if shared.CheckError(err) {
		log.Fatal("Couldnt open the database!")
	}

	database = d

	//create table
	res, err := database.Exec("CREATE TABLE IF NOT EXISTS media_info (author TEXT, tweet_id INTEGER, media_hash TEXT UNIQUE, media_url TEXT UNIQUE);")
	if shared.CheckError(err) {
		log.Fatal(res)
	}

	fmt.Println("Database opened", database_path)
}

//saves a media info struct on the database
func SaveMediaInfo(info MediaInfo) {
	_, err := database.Exec(fmt.Sprintf("INSERT INTO media_info(author, tweet_id, media_hash, media_url) VALUES ('%s', %d, '%s', '%s');", info.Author, info.TweetID, info.MediaHash, info.MediaURL))
	fmt.Println("Saved new media info:", info)
	shared.CheckError(err)
}

//check if a hash is already on the database
func AssetAlreadyUploaded(hash string, url string) bool {
	fmt.Println("Checking for hash", hash, "and url", url, "on database")
	rows, err := database.Query("SELECT media_hash, media_url FROM media_info;")
	if shared.CheckError(err) {
		return false
	}

	defer rows.Close()
	for rows.Next() {
		var stored_hash, stored_url string
		rows.Scan(&stored_hash, &stored_url)

		if hash == stored_hash || url == stored_url {
			fmt.Println("Database match found, aborting upload")
			return true
		}
	}

	fmt.Println("Coudln't find a match, proceding!")
	return false
}

//closes the database connection
func Close() {
	database.Close()
	fmt.Println("Database closed")
}
