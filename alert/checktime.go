package alert

import (
	"fmt"
	"migrai/gameinfo"
	"migrai/message"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Check1 bool
var Check2 bool
var Check3 bool
var Check4 bool
var Check5 bool
var Check6 bool

var Turn int

func CheckTime(s *discordgo.Session, r *discordgo.Ready) {

	for {

		currentTurn := gameinfo.GetTurn()
		fmt.Printf("turn: %d, currentturn: %d\n", Turn, currentTurn)
		if Turn != currentTurn {
			Check6 = false
			Check5 = false
			Check4 = false
			Check3 = false
			Check2 = false
			Check1 = false

			Turn = currentTurn
		}

		sb := gameinfo.GetTime()

		parsedTime := gameinfo.ParseTime(sb)

		if parsedTime.Hours == 47 && !Check1 {
			Check1 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 23 && !Check2 {
			Check2 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 11 && !Check3 {
			Check3 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 29 && !Check4 {
			Check4 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 14 && !Check5 {
			Check5 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 4 && !Check6 {
			Check6 = true
			go message.Send(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes < 3 {
			Check6 = false
			Check5 = false
			Check4 = false
			Check3 = false
			Check2 = false
			Check1 = false
		}

		time.Sleep(1 * time.Second)
	}

}
