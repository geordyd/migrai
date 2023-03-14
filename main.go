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
		log.Fatal(err)
		return -1
	}

	sb := string(body)

	turnStripped := strings.ReplaceAll(sb, "beurt ", "")

	turnInt, _ := strconv.Atoi(turnStripped)

	return turnInt

}

func checkTime(s *discordgo.Session, r *discordgo.Ready) {

	for {

		currentTurn := getTurn()

		if turn != currentTurn {
			check6 = false
			check5 = false
			check4 = false
			check3 = false
			check2 = false
			check1 = false

			turn = currentTurn
		}

		resp, err := http.Get("http://localhost:3000/time")
		if err != nil {
			fmt.Println(err)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		sb := string(body)

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
			check5 = true
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

	hours, _ := strconv.Atoi(splittedTime[0])
	min, _ := strconv.Atoi(splittedTime[1])
	sec, _ := strconv.Atoi(splittedTime[2])

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

		resp, err := http.Get("http://localhost:3000/time")
		if err != nil {
			fmt.Println(err)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		sb := string(body)

		_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Turn %d, time left: %s", turn, sb))
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
	if strings.HasPrefix(m.Content, "!african-facts random") {
		//reference to the africanfact file

		_, err := s.ChannelMessageSend(m.ChannelID, " ")

		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.HasPrefix(m.Content, "!help") {

		_, err := s.ChannelMessageSend(m.ChannelID, "```Available commands:\n - !help\n - !timeleft\n - !link\n - !changelog - !link\n - !!african-facts random```")
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
