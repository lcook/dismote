package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (c *Commands) Help() {
	fields := []*discordgo.MessageEmbedField{}

	for k := range c.FuncMap {
		if k == "default" {
			continue
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name: c.Prefix + k, Value: "\u200e", Inline: true,
		})
	}

	embed := discordgo.MessageEmbed{
		Title:     fmt.Sprintf("__%s's commands__", c.Session.State.User.Username),
		Fields:    fields,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: c.Session.State.User.AvatarURL("2048")},
		Color:     randomColor(),
	}

	err := c.Session.ChannelMessageDelete(c.Message.ChannelID, c.Message.ID)
	if err != nil {
		log.Println("Help:", err)
	}

	_, err = c.Session.ChannelMessageSendEmbed(c.Message.ChannelID, &embed)
	if err != nil {
		log.Println("Help:", err)
	}
}
