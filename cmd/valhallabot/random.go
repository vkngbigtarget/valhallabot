package main

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/bwmarrin/discordgo"
	"github.com/hink/valhallabot/cmd/valhallabot/pkg/models"
)

func countImakPinged(s *discordgo.Session, m *discordgo.MessageCreate) {
	val := db.Get(BucketCounters, []byte("imakPings"))
	now := time.Now()
	data := new(models.CountData)
	if val != nil {
		if err := json.Unmarshal(val, data); err != nil {
			log.Error(err.Error())
			return
		}
	} else {
		data.Count = 1
	}

	data.Count = data.Count + 1
	data.LastUpdated = now.Unix()
	jData, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
		return
	}
	err = db.Save(BucketCounters, []byte("imakPings"), jData)
	if err != nil {
		log.Error(err.Error())
		return
	}

	rMsgs := []string{
		"Really? That's a total of #pings.",
		"It never get's old, Imak has been pinged #pings times.",
		"He's already been pinged #pings times. I bet you won't do it again.",
		"When will you learn? Once is not enough! Neither are the previous #pings.",
		"We need to get those numbers up. Those are rookie numbers. That's a total of only #pings pings.",
		"Why you have to be mad? It is only ping? #pings total.",
		"Rally? ... nah, PING. That's #pings times.",
		"I doubt he'll respond. He's already been pinged #pings times.",
		"Tha ping goes skrraaaa, pop, pop, ka ka ka. That's #pings now.",
		"2 + 2 is 4, plus one ping that's #pings. Quick maths.",
		"Hello? Is it #pings pings you're looking for?",
		"Never gonna give you up. Never gonna let you down. Gonna ping you #pings times and notify you!",
		"#yes-yes-#pings!",
		"You have #pings unread messages!",
		"Ding! #pings aren't done! Ding! #pings aren't done!",
		"Oh mama mia, mama mia! Mama mia #pings and more to go!",
		"#pings reeeeeeeeeeeeeeeeeeeeeeee's recorded. Autism levels still rising...",
		"One, two, three, four, #pings. Everybody in the car, so come on, let's ride!",
		"These pings are THICC with #pings c's",
		"#pings dollar sucky sucky! Imak love you long time!",
		"Whoa! Imak ate a total of #pings :eggplant:s!",
		"If Imak had a nickel for everytime he's been pinged, he'd have #pings nickels. That's a lot!",
		"Adjusting for inflation, I bet Imak could afford a house with $#pings. Or at least part of one.",
		"Who left the light on in Imak's office? That's #pings times since he moved in!",
		"If #pings trees fall in the forest and Imak isn't around to hear them, do they make a noise?",
		"On the first day of Christmas, my true love sent to me: #pings pings to Imak!",
		"Real estate? Nah. Ping estate. That's #pings pings, already!",
		"Maybe he didn't catch the last #pings pings. Better ping him, again.",
		"#pings :o:s... Er pings!",
		"Another day. Another ping. #pings pings to be exact.",
		"Does Imak ever think how much one person can get pinged in one day? Because #pings is a lot.",
		":ping: is what Imak looks like in real life. This has been confirmed by #pings witnesses.",
		"Pinging Imak #pings times is the new spice of life.",
		"If the combined energy of #pings pings could create heat, Imak would be sooooo :fire:!",
		"Is this the real life? / Is this just fantasy? / Caught in a landslide / Of #pings pings",
		"Pinging Imak #pings times is more popular than Florida Georgia Line and Nickelback combined.",
		"Pinging Imak is so therapeutic. #pings doctors approve.",
		"Imak has brought so many of us together. We just want to thank him #pings times.",
		"To ping or not to ping. That is the question. Who are we kidding? #pings total pings!",
		"You burn #pings calories just by pinging Imak. Who knew?",
		"Pings on pings on pings on pings... #pings pings.",
		"#pings doctors recommend diet, exercise, and a crap ton of pings to Imak.",
		"Ping ya later, bishes! #pings pings.",
		"PINGGG! PINGGG! That's the sound a button on the pants of a #pings pound man makes ricocheting off the walls.",
		"Only #pings? I think Imak needs more pings in his life.",
		"Imak teamkilled #pings total players. Someone needs more disciprine.",
		"I'm running out of catchy things to say. #pings total pings.",
		"If Imak had an :eggplant: for every ping, he'd have #pings total and a really good time.",
		"We've pinged Imak #pings times and it's time to stop. Change my mind.",
		"Hi. Yea. I would like to order a venti Latte Macchiato with #pings pumps, please.",
		"Whoa! Did you see that? A stampede of #pings pings just flew by!",
	}

	rand.Seed(time.Now().Unix())
	msg := rMsgs[rand.Intn(len(rMsgs))]
	msg = strings.Replace(msg, "#pings", strconv.Itoa(data.Count), -1)
	s.ChannelMessageSend(m.ChannelID, msg)
}
