package main

import (
	"fmt"
	"github.com/NeedMoreVolume/shootyhoops/handlers"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var botUser *discordgo.User

func main() {

	token := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: " + err.Error())
		return
	}

	dg.AddHandler(messageHandler)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	botUser = dg.State.User

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("shootyhoops is now raining buckets.  Press CTRL-C to make it do the epic Rockets-Warriors meltdown and go ice cold from downtown.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	var botMention bool
	for _, mention := range m.Mentions {
		if mention.Mention() == botUser.Mention() {
			botMention = true
		}
	}
	if !botMention {
		return
	}
	if strings.HasPrefix(m.Content, botUser.Mention()) {
		m.Content = strings.Replace(m.Content, botUser.Mention(), "", 1)
	}
	botMention2 := strings.Replace(botUser.Mention(), "@", "@!", 1) + " "
	if strings.HasPrefix(m.Content, botMention2) {
		m.Content = strings.Replace(m.Content, botMention2, "", 1)
	}

	handlers.BaseHandler(s, m)
	return
}
