package commands

import (
	"log"
)

func (c *Commands) Clear() {
	err := c.Session.ChannelMessageDelete(c.Message.ChannelID, c.Message.ID)
	if err != nil {
		log.Println("Clear:", err)
		return
	}

	ids := []string{}

	batch, _ := c.Session.ChannelMessages(c.Message.ChannelID, 100, "", "", "")
	for _, message := range batch {
		ids = append(ids, message.ID)
	}

	err = c.Session.ChannelMessagesBulkDelete(c.Message.ChannelID, ids)
	if err != nil {
		log.Println("Clear:", err)
		return
	}
}
