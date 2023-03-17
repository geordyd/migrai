package main

import (
	"flag"
	"fmt"
	"migrai/alert"
	"migrai/message"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {

		alert.Turn = -1

		alert.Check1 = false
		alert.Check2 = false
		alert.Check3 = false
		alert.Check4 = false
		alert.Check5 = false
		alert.Check6 = false

		go alert.CheckTime(s, r)
	})

	dg.AddHandler(message.Create)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
