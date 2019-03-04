package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetValhallaStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := ""
	server, err := bm.Server(1155175)
	if err != nil {
		msg = "Server Error"
	} else {
		poppedMsg := "Get in there and help pop!"
		if server.Attributes.Players >= 40 {
			poppedMsg = "We're popped boys!"
		}
		msg = fmt.Sprintf("There are currently %d players in Squad on Valhalla. %s", server.Attributes.Players, poppedMsg)
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}
