package gameinfo

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetTurn() int {
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
