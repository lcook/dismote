package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) Listen() {
	status := strings.TrimPrefix(h.Message.Content, h.Config.Prefix+"listen ")

	embed := discordgo.MessageEmbed{
		Title: "__Status updated__!",
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Now listening to " + status, Value: "\u200e", Inline: false},
		},
		Color: randomColor(),
	}

	err := h.Session.UpdateListeningStatus(status)
	if err != nil {
		log.Println("Listen:", err)
	}

	err = h.Session.ChannelMessageDelete(h.Message.ChannelID, h.Message.ID)
	if err != nil {
		log.Println("Listen:", err)
	}

	_, err = h.Session.ChannelMessageSendEmbed(h.Message.ChannelID, &embed)
	if err != nil {
		log.Println("Listen:", err)
	}
}
