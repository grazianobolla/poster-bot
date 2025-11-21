package discord

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"shitposter-bot/shared"
	"shitposter-bot/tenor"
	"shitposter-bot/uploader"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// utf8 emojis used to comunicate with the bot
const UPLOAD_EMOJI = "\xE2\xAC\x86"
const CHECK_MARK_EMOJI = "\xE2\x9C\x85"
const ALARM_CLOCK_EMOJI = "\xE2\x8F\xB0"
const ERROR_EMOJI = "\xE2\x9D\x8C"

// keyword used to prevent an asset from being marked ass uploadable
const CANCEL_KEYWORD = "ANTIFUNA"

// discord client connection
var client *discordgo.Session

// starts the listening
func start_connection(token string) {
	var err error
	client, err = discordgo.New("Bot " + token)
	shared.CheckError(err)

	client.AddHandler(guild_create)
	client.AddHandler(message_create)
	client.AddHandler(message_reaction_add)

	//open the websocket and start listening
	err = client.Open()
	if shared.CheckError(err) {
		log.Fatal("Can't start Discord connection")
		return
	}

	fmt.Println("Shitposter Bot Discord is now up and running")
}

// called when the bot connects to a guild
func guild_create(s *discordgo.Session, m *discordgo.GuildCreate) {
	fmt.Printf("Connected to server %s\n", m.Name)
}

// called when a client sends a message
func message_create(s *discordgo.Session, m *discordgo.MessageCreate) {
	//check if the author of the message is an user
	if m.Author.Bot {
		return
	}

	//check if the cancel keyword has been typed
	for _, s := range strings.Fields(m.Content) {
		if s == CANCEL_KEYWORD {
			return
		}
	}

	if url, text, ok := get_media_url_text(m); ok {
		fmt.Println(url)
		mime := shared.GetContentType(url)

		if !strings.Contains(mime, "image") && !strings.Contains(mime, "video") {
			return
		}

		if !add_reaction(m.ChannelID, m.ID, UPLOAD_EMOJI) {
			return
		}

		temp_asset := Asset{
			MessageID:    m.ID,
			ChannelID:    m.ChannelID,
			AuthorName:   get_username(m.Author.ID),
			CreationTime: time.Now(),
			Uploaded:     false,
			Url:          url,
			Text:         parse_text(text),
		}

		assets = append(assets, temp_asset)
		fmt.Println("Added asset url", url, text)
	}
}

// called when a message is given a reaction
func message_reaction_add(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.UserID == s.State.User.ID {
		return
	}

	//if we clicked the upload emoji, check if the asset exists and then upload
	if m.Emoji.Name == UPLOAD_EMOJI {
		if is_asset, asset := is_asset(m.MessageID); is_asset && !asset.Uploaded {
			asset.Uploaded = true

			if uploader.UploadAsset(asset.AuthorName, asset.Text, asset.Url) {
				add_reaction(m.ChannelID, m.MessageID, CHECK_MARK_EMOJI)
			} else {
				add_reaction(m.ChannelID, m.MessageID, ERROR_EMOJI)
			}
		}
	}
}

// turns the discord text, wich might containt weird emoji encoding or user data to a human readable format
func parse_text(src string) string {
	reg := regexp.MustCompile(`<[@|:](.*?)>`)
	src = reg.ReplaceAllString(src, "")
	return src
}

// returns the media url from a discord message
func get_media_url_text(m *discordgo.MessageCreate) (string, string, bool) {
	msg_fields := strings.Fields(m.Content)
	msg_len := len(msg_fields)
	//for files
	if len(m.Attachments) > 0 {
		return m.Attachments[0].URL, m.Content, true
	}

	//for tenor gifs
	if msg_len > 0 {
		if strings.Contains(msg_fields[0], "https://tenor.com") {
			if url, ok := tenor.GetGIFbyURL(msg_fields[0]); ok {
				text := strings.Join(msg_fields[0:], " ")

				if msg_len > 1 {
					text = strings.Join(msg_fields[1:], " ")
				}
				return url, text, true
			}
		}
	}

	//for urls
	for i, field := range msg_fields {
		resp, err := http.Get(field)
		if err == nil && resp.StatusCode == http.StatusOK {
			return field, strings.Join(msg_fields[i+1:], " "), true
		}
	}

	return "", "", false
}

// returns the username of a userid
func get_username(userid string) string {
	user, err := client.User(userid)

	if shared.CheckError(err) {
		return ""
	}

	return user.Username
}

// safe function for adding reactions to messages
func add_reaction(channelID string, messageID string, emoji string) bool {
	m, err := client.ChannelMessage(channelID, messageID)
	if err == nil {
		client.MessageReactionAdd(m.ChannelID, m.ID, emoji)
		return true
	}

	return false
}
