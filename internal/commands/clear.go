package commands

import (
	"log"
)

func (h *Handler) Clear() {
	err := h.Session.ChannelMessageDelete(h.Message.ChannelID, h.Message.ID)
	if err != nil {
		log.Println("Clear:", err)
		return
	}

	ids := []string{}

	batch, _ := h.Session.ChannelMessages(h.Message.ChannelID, 100, "", "", "")
	for _, message := range batch {
		ids = append(ids, message.ID)
	}

	err = h.Session.ChannelMessagesBulkDelete(h.Message.ChannelID, ids)
	if err != nil {
		log.Println("Clear:", err)
		return
	}
}
