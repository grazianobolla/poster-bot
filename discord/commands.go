package discord

import (
	"github.com/bwmarrin/discordgo"
)

var (
	cmd_array = []*discordgo.ApplicationCommand{
		{
			Name:        "basic",
			Description: "A basic, basic command.",
		},
	}

	cmd_handler_array = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"penis": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just basic!",
				},
			})
		},
	}
)
