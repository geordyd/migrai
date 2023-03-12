package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type currentTime struct {
	Hours   int
	Minutes int
	Seconds int
}

var token string

var turn int

var check1 bool
var check2 bool
var check3 bool
var check4 bool
var check5 bool
var check6 bool

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

		turn = -1

		check1 = false
		check2 = false
		check3 = false
		check4 = false
		check5 = false
		check6 = false

		go checkTime(s, r)
	})

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

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

func getTurn() int {
	resp, err := http.Get("http://localhost:3000/turn")
	if err != nil {
		fmt.Println(err)
		return -1
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	sb := string(body)

	turnStripped := strings.ReplaceAll(sb, "beurt ", "")

	turnInt, _ := strconv.Atoi(turnStripped)

	return turnInt

}

func getTime() string {
	resp, err := http.Get("http://localhost:3000/time")
	if err != nil {
		fmt.Println(err)
		return "69:69:69"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "69:69:69"
	}

	return string(body)
}

func checkTime(s *discordgo.Session, r *discordgo.Ready) {

	for {

		currentTurn := getTurn()
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

		sb := getTime()

		parsedTime := parseTime(sb)

		if parsedTime.Hours == 47 && !check1 {
			check1 = true
			go sendMessage(s, r, sb)
		} else if parsedTime.Hours == 23 && !check2 {
			check2 = true
			go sendMessage(s, r, sb)
		} else if parsedTime.Hours == 11 && !check3 {
			check3 = true
			go sendMessage(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 29 && !check4 {
			check4 = true
			go sendMessage(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 14 && !check5 {
			check5 = true
			go sendMessage(s, r, sb)
		} else if parsedTime.Hours == 0 && parsedTime.Minutes == 4 && !check6 {
			check6 = true
			go sendMessage(s, r, sb)
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

func parseTime(sb string) currentTime {

	splittedTime := strings.Split(sb, ":")

	if len(splittedTime) == 2 {
		splittedTime = append([]string{"00"}, splittedTime...)
	} else if len(splittedTime) == 1 {
		splittedTime = append([]string{"00", "00"}, splittedTime...)
	}

	hours, err := strconv.Atoi(splittedTime[0])
	if err != nil {
		fmt.Println(err)
	}
	min, err := strconv.Atoi(splittedTime[1])
	if err != nil {
		fmt.Println(err)
	}
	sec, err := strconv.Atoi(splittedTime[2])
	if err != nil {
		fmt.Println(err)
	}

	time := currentTime{
		Hours:   hours,
		Minutes: min,
		Seconds: sec,
	}

	return time

}

func sendMessage(s *discordgo.Session, m *discordgo.Ready, time string) {
	_, err := s.ChannelMessageSend("1080954466896982027", fmt.Sprintf("@everyone Turn %d, time left: %s\nhttp://atwar-game.com/games/?link=7878030246", turn, time))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!timeleft") {

		sb := getTime()

		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Turn %d, time left: %s", turn, sb))
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

		changelog := readChangelog()

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

func readChangelog() string {
	content, err := os.ReadFile("changelog.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
