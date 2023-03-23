package waifu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Waifu struct {
	Images []Images `json:"images"`
}

type Images struct {
	URL string `json:"url"`
}

func Get() string {

	resp, err := http.Get("https://api.waifu.im/search/?included_tags=waifu")
	if err != nil {
		fmt.Println(err)
		return "No waifu :("
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "No waifu :("
	}

	var result Waifu
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON")
		return "No waifu :("
	}

	return result.Images[0].URL
}
