package main

import "github.com/bwmarrin/discordgo"

func ImakPingCounter(s *discordgo.Session, m *discordgo.MessageCreate) {
	val := db.Get(BucketCounters, []byte("imakPings"))
	if val == nil {
		count = 1
	} else {

	}
}
