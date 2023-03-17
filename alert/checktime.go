package alert

import (
	"fmt"
	"migrai/gameinfo"
	"migrai/message"
	"time"

	"github.com/bwmarrin/discordgo"
)

func CheckTime(s *discordgo.Session, r *discordgo.Ready) {

	check1 := false
	check2 := false
	check3 := false
	check4 := false
	check5 := false
	check6 := false

	turn := -1

	for {

		currentTurn := gameinfo.GetTurn()
		fmt.Printf("turn: %d, currentturn: %d\n", turn, currentTurn)
		if turn != currentTurn {
			check6 = false
			check5 = false
			check4 = false
			check3 = false
			check2 = false
			check1 = false

			turn = currentTurn
		}

		sb := gameinfo.GetTime()

		parsedTime := gameinfo.ParseTime(sb)

		if parsedTime.Hours == 47 && !check1 {
			check1 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 23 && !check2 {
			check2 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 11 && !check3 {
			check3 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 29 && !check4 {
			check4 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 14 && !check5 {
			check5 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 4 && !check6 {
			check6 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes < 3 {
			check6 = false
			check5 = false
			check4 = false
			check3 = false
			check2 = false
			check1 = false
		}

		time.Sleep(1 * time.Second)
	}

}
