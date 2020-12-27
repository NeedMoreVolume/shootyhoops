package handlers

import (
	"fmt"
	"github.com/NeedMoreVolume/shootyhoops/handlers/help"
	"github.com/NeedMoreVolume/shootyhoops/handlers/nba"
	"github.com/NeedMoreVolume/shootyhoops/handlers/ncaa"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func BaseHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := m.Content
	response := "unrecognized request..."
	switch {
	case strings.Contains(m.Content, "help"), strings.Contains(m.Content, "halp"), strings.Contains(m.Content, "lifealert"):
		response = help.Help()
	case strings.HasPrefix(message, "nba-games"):
		response = nba.Games(message)
	case strings.HasPrefix(message, "nba-standings"):
		response = nba.Standings(message)
	case strings.HasPrefix(message, "ncaa-games"):
		response = ncaa.Games(message)
	case strings.HasPrefix(message, "ncaa-standings"):
		response = ncaa.Standings(message)
	}

	if len(response) > 2000 {
		BulkMessageHandler(s, m, response)
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, response)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func BulkMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate, message string) {
	split := strings.Split(message, "```")
	var partialResponse string
	for i, game := range split {
		if i != 0 && (i%10) == 0 {
			// send response buffer then reset it
			_, err := s.ChannelMessageSend(m.ChannelID, partialResponse)
			if err != nil {
				fmt.Println("index: ", i)
				fmt.Println(err)
			}
			partialResponse = ""
		}
		// if this is nothing, skip it
		if game == "" || game == " " {
			continue
		}
		// add current game to response buffer
		partialResponse += "```" + game + "```"
	}
	// send remaining chunk
	_, err := s.ChannelMessageSend(m.ChannelID, partialResponse)
	if err != nil {
		fmt.Println(err)
	}
	return
}
