package main

import (
	"os"
	"os/signal"
	"syscall"

	apidb "github.com/hink/apidb/bolt"
	"github.com/hink/valhallabot/pkg/battlemetrics"

	"github.com/bwmarrin/discordgo"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/hink/valhallabot/internal/pkg/config"
)

var db *apidb.Database
var cfg *config.Config
var discord *discordgo.Session
var bm *battlemetrics.Client

func main() {
	var err error
	// log setup
	log.SetHandler(text.New(os.Stderr))

	// load config
	cfg, err = config.Load("valhallabot.cfg")
	if err != nil {
		log.Fatal(err.Error())
	}

	// load database
	db, err = apidb.Open(cfg.Database.Path, databaseBuckets)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// discord
	discord, err = discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {
		log.Fatal(err.Error())
	}

	// battlemetrics
	bm = battlemetrics.New(cfg.BattleMetrics.Token)

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		log.Fatalf("error opening connection, %s", err.Error())
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info("ValhallaBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
