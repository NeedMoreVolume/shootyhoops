package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"shootyhoops/handlers/avatar"
	"shootyhoops/handlers/help"
	"shootyhoops/handlers/nba"
	"shootyhoops/handlers/ncaa"
	"strings"
)

type Handler struct {
	bot *discordgo.User
}

func NewHandler(b *discordgo.User) *Handler {
	handler := &Handler{bot: b}
	return handler
}

func (h *Handler) BaseHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if h.isBotAuthor(m.Author) {
		return
	}

	if !h.isBotMention(m.Mentions) {
		return
	}

	m.Content = h.stripBotMention(m.Content)

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
	case strings.HasPrefix(message, "set-avatar"):
		err := avatar.SetAvatar(s, message)
		if err != nil {
			fmt.Println(err.Error())
			_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
		}
		return
	}

	if len(response) > 2000 {
		h.sendBulkResponse(s, m, response)
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, response)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (h *Handler) sendBulkResponse(s *discordgo.Session, m *discordgo.MessageCreate, message string) {
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
	if partialResponse != "" {
		_, err := s.ChannelMessageSend(m.ChannelID, partialResponse)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}

func (h *Handler) isBotAuthor(author *discordgo.User) bool {
	return author.ID == h.bot.ID
}

func (h *Handler) isBotMention(mentions []*discordgo.User) bool {
	for _, mention := range mentions {
		if mention.Mention() == h.bot.Mention() {
			return true
		}
	}
	return false
}

func (h *Handler) stripBotMention(message string) string {
	if strings.HasPrefix(message, h.bot.Mention()) {
		message = strings.Replace(message, h.bot.Mention(), "", 1)
	}
	botMention2 := strings.Replace(h.bot.Mention(), "@", "@!", 1) + " "
	if strings.HasPrefix(message, botMention2) {
		message = strings.Replace(message, botMention2, "", 1)
	}
	return message
}
