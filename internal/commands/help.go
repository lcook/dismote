package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) Help() {
	fields := []*discordgo.MessageEmbedField{}

	for k := range h.CmdMap {
		if k == "default" {
			continue
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name: h.Prefix + k, Value: "\u200e", Inline: true,
		})
	}

	embed := discordgo.MessageEmbed{
		Title:     fmt.Sprintf("__%s's commands__", h.Session.State.User.Username),
		Fields:    fields,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: h.Session.State.User.AvatarURL("2048")},
		Color:     randomColor(),
	}

	err := h.Session.ChannelMessageDelete(h.Message.ChannelID, h.Message.ID)
	if err != nil {
		log.Println("Help:", err)
	}

	_, err = h.Session.ChannelMessageSendEmbed(h.Message.ChannelID, &embed)
	if err != nil {
		log.Println("Help:", err)
	}
}
