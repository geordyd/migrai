package gameinfo

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type currentTime struct {
	Hours   int
	Minutes int
	Seconds int
}

func GetTime() string {
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

func ParseTime(sb string) currentTime {

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
