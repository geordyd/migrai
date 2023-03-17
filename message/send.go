package message

import (
	"fmt"
	"migrai/gameinfo"

	"github.com/bwmarrin/discordgo"
)

func Send(s *discordgo.Session, m *discordgo.Ready, time string) {
	currentTurn := gameinfo.GetTurn()
	_, err := s.ChannelMessageSend("1080954466896982027", fmt.Sprintf("@everyone Turn %d, time left: %s\nhttp://atwar-game.com/games/?link=7878030246", currentTurn, time))
	if err != nil {
		fmt.Println(err)
		return
	}
}
