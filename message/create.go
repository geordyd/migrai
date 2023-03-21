package message

import (
	"fmt"
	"migrai/changelog"
	"migrai/gameinfo"
	"migrai/mugshot"
	"migrai/waifu"
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

		_, err := s.ChannelMessageSend(m.ChannelID, "```Available commands:\n - !help\n - !timeleft\n - !link\n - !changelog\n - !waifu```")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.HasPrefix(m.Content, "!waifu") {

		username := m.Author.Username

		waifuUrl := waifu.Get()

		message := fmt.Sprintf("Congratulations %s on finding your waifu!", username)

		embeddedMessage := &discordgo.MessageEmbed{
			Image: &discordgo.MessageEmbedImage{
				URL: waifuUrl,
			},
			Description: message,
		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embeddedMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
	if strings.HasPrefix(m.Content, "!mugshot") {

		images, mugshotData, err := mugshot.Get()
		if err != nil {
			fmt.Println("Cannot get mugshot")
			return
		}

		username := m.Author.Username

		mugshotDataText := fmt.Sprintf("Sex: %s\nHeight: %s\nWeight: %s\nHair: %s\nEyes: %s\nRace: %s\nSex Offender: %s\nOffense: %s", mugshotData.Sex, mugshotData.Height, mugshotData.Weight, mugshotData.Hair, mugshotData.Eyes, mugshotData.Race, mugshotData.SexOffender, mugshotData.Offense)

		message := fmt.Sprintf("Congratulations %s this is your mugshot!\n%s", username, mugshotDataText)

		files := []*discordgo.File{}

		for i, image := range images {
			files = append(files, &discordgo.File{
				Name:   fmt.Sprintf("%s%d.png", mugshotData.ID, i),
				Reader: image,
			})
		}

		embeddedMessage := &discordgo.MessageSend{
			Content: message,
			Files:   files,
		}

		_, err = s.ChannelMessageSendComplex(m.ChannelID, embeddedMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
