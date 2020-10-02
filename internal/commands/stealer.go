package commands

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	EmojiPrefix   string = "<"
	EmojiSuffix   string = ">"
	EmojiAnimated string = "a:"
	EmojiIcon     string = ":"
)

type Emoji struct {
	ID   string
	Name string
	Type string
	Data string
}

func (h *Handler) Stealer() {
	if !strings.HasPrefix(h.Message.Content, EmojiPrefix) {
		return
	}

	emojis := []Emoji{}

	content := strings.Split(strings.Replace(h.Message.Content, " ", "", -1), EmojiSuffix)
	for _, element := range content {
		if element == "" || !strings.HasPrefix(element, EmojiPrefix) {
			continue
		}

		element = cleanPrefix(element)

		emoji := Emoji{
			ID:   strings.Split(element, ":")[1],
			Name: strings.Split(element, ":")[0],
		}

		resp, err := http.Get(endpointEmoji(emoji.ID))
		if err != nil {
			err, _ := h.Session.ChannelMessageSend(h.Message.ChannelID, "Ooops! I was unable to fetch "+formatEmoji(emoji))
			if err != nil {
				log.Println("Stealer:", err)
			}

			continue
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		emoji.Type = http.DetectContentType(body)
		emoji.Data = "data:" + emoji.Type + ";base64," + base64.StdEncoding.EncodeToString(body)

		emj, err := h.Session.GuildEmojiCreate(h.Message.GuildID, emoji.Name, emoji.Data, nil)
		if err != nil {
			message := fmt.Sprintf("Ooops! I was unable to add %s, %s", formatEmoji(emoji), err)
			err, _ := h.Session.ChannelMessageSend(h.Message.ChannelID, message)

			if err != nil {
				log.Println("Stealer:", err)
			}

			continue
		}

		emoji.ID = emj.ID
		emoji.Name = emj.Name

		emojis = append(emojis, emoji)
	}

	err := h.Session.ChannelMessageDelete(h.Message.ChannelID, h.Message.ID)
	if err != nil {
		log.Println("Stealer:", err)
	}

	var message string
	for _, emoji := range emojis {
		message += formatEmoji(emoji)
	}

	embed := discordgo.MessageEmbed{
		Title: "__Emotes added__!",
		Fields: []*discordgo.MessageEmbedField{
			{Name: message, Value: "\u200e", Inline: false},
		},
		Color: randomColor(),
	}

	_, err = h.Session.ChannelMessageSendEmbed(h.Message.ChannelID, &embed)
	if err != nil {
		log.Println("Stealer:", err)
	}
}

func formatEmoji(e Emoji) string {
	var str string

	switch e.Type {
	case "image/png":
		str = fmtEmojiIcon(e)
	case "image/gif":
		str = fmtEmojiAnimated(e)
	}

	return str
}

func cleanPrefix(s string) string {
	str := strings.TrimPrefix(s, EmojiPrefix)

	str = strings.TrimPrefix(str, EmojiIcon)
	str = strings.TrimPrefix(str, EmojiAnimated)

	return str
}

func fmtEmojiIcon(emj Emoji) string {
	return EmojiPrefix + EmojiIcon + emj.Name + ":" + emj.ID + EmojiSuffix
}

func fmtEmojiAnimated(emj Emoji) string {
	return EmojiPrefix + EmojiAnimated + emj.Name + ":" + emj.ID + EmojiSuffix
}

func endpointEmoji(id string) string {
	return discordgo.EndpointCDN + "emojis/" + id
}
