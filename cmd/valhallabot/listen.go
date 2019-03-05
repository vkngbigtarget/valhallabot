package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
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
	default:
		// imak pings
		if strings.Contains(m.Content, UserImak) {
			countImakPinged(s, m)
		}
	}
}
