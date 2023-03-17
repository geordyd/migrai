package message

import (
	"fmt"
	"migrai/changelog"
	"migrai/gameinfo"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Create(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!timeleft") {

		sb := gameinfo.GetTime()
		currentTurn := gameinfo.GetTurn()

		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Turn %d, time left: %s", currentTurn, sb))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.HasPrefix(m.Content, "!link") {

		_, err := s.ChannelMessageSend(m.ChannelID, "http://atwar-game.com/games/?link=7878030246")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.HasPrefix(m.Content, "!changelog") {

		changelog := changelog.Read()

		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("```%s```", changelog))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.HasPrefix(m.Content, "!help") {

		_, err := s.ChannelMessageSend(m.ChannelID, "```Available commands:\n - !help\n - !timeleft\n - !link\n - !changelog```")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
