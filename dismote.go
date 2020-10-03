package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/lcook/dismote/internal/commands"
	"gitlab.com/lcook/dismote/internal/config"
)

func main() {
	var Config string

	flag.StringVar(&Config, "c", "config.yaml", "Configuration file")
	flag.Parse()
LoadConfig:
	configFmt := "(\"" + Config + "\")"

	configFile, err := config.LoadConfig(Config)
	if err != nil {
		log.Fatalln("Error loading configuration file", err)
	}

	log.Println("Loaded configuration file", configFmt)

	if len(configFile.Channels) < 1 {
		log.Fatalln("Error: No channels provided in configuration file", configFmt)
	}

	dg, err := discordgo.New(configFile.Token)
	if err != nil {
		log.Fatalln("Error creating Discord session,", err)
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if configFile.Bot && m.Author.ID == s.State.User.ID || !configFile.IsChannel(m.ChannelID) {
			return
		}

		if !configFile.Bot && configFile.Owner == "" {
			configFile.Owner = s.State.User.ID
		}

		cmds := commands.New(s, m, &configFile)

		cmds.Register(commands.Modules{
			"default": {cmds.Stealer, commands.PermAll},
			"info":    {cmds.Info, commands.PermAll},
			"clear":   {cmds.Clear, commands.PermOwner},
			"help":    {cmds.Help, commands.PermAll},
			"listen":  {cmds.Listen, commands.PermOwner},
		})

		cmds.Execute(m.Content)
	})

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		log.Fatalln("Error opening connection,", err)
	}

	log.Println("Successfully started Discord session")
	log.Println("Now listening on", len(configFile.Channels), "channel(s)")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)

	for s := range ch {
		log.Println("Receieved signal:", s)

		abort := false

		switch s {
		case syscall.SIGHUP:
			log.Println("Reloading settings...")
			goto LoadConfig
		case syscall.SIGTERM:
			fallthrough
		case os.Interrupt:
			log.Println("Application shutting down...")

			abort = true
		}

		if abort {
			break
		}
	}

	dg.Close()
}
