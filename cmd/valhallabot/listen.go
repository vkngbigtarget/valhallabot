package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// listen for new messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "!server":
		GetValhallaStatus(s, m)
	case
	default:
		if strings.Contains(m.Content, "@VKNG | imak") {

		}
	}
}
