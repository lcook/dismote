package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) Info() {
	guild, _ := h.Session.Guild(h.Message.GuildID)
	emojis := guild.Emojis
	emojiIcon := 0
	emojiAnimated := 0

	for _, emoji := range emojis {
		if emoji.Animated {
			emojiAnimated++
		} else {
			emojiIcon++
		}
	}

	embed := discordgo.MessageEmbed{
		Fields: []*discordgo.MessageEmbedField{
			{Name: fmt.Sprintf("__%s server statistics__", guild.Name), Value: "\u200e", Inline: false},
			{Name: "**ID**", Value: guild.ID, Inline: false},
			{Name: "**Roles**", Value: strconv.Itoa(len(guild.Roles)), Inline: false},
			{Name: "**Members**", Value: strconv.Itoa(guild.ApproximateMemberCount), Inline: false},
			{Name: "**Channels**", Value: strconv.Itoa(len(guild.Channels)), Inline: false},
			{Name: "**Emotes**", Value: fmt.Sprintf("Icons: **%d**, Animated: **%d**", emojiIcon, emojiAnimated), Inline: false},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: guild.IconURL()},
		Color:     randomColor(),
	}

	err := h.Session.ChannelMessageDelete(h.Message.ChannelID, h.Message.ID)
	if err != nil {
		log.Println("Info:", err)
	}

	_, err = h.Session.ChannelMessageSendEmbed(h.Message.ChannelID, &embed)
	if err != nil {
		log.Println("Info:", err)
	}
}
