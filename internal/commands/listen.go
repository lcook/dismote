package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (c *Commands) Listen() {
	status := strings.TrimPrefix(c.Message.Content, c.Prefix+"listen ")

	embed := discordgo.MessageEmbed{
		Title: "__Status updated__!",
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Now listening to " + status, Value: "\u200e", Inline: false},
		},
		Color: randomColor(),
	}

	err := c.Session.UpdateListeningStatus(status)
	if err != nil {
		log.Println("Listen:", err)
	}

	err = c.Session.ChannelMessageDelete(c.Message.ChannelID, c.Message.ID)
	if err != nil {
		log.Println("Listen:", err)
	}

	_, err = c.Session.ChannelMessageSendEmbed(c.Message.ChannelID, &embed)
	if err != nil {
		log.Println("Listen:", err)
	}
}
