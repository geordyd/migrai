package changelog

import (
	"log"
	"os"
)

func Read() string {
	content, err := os.ReadFile("changelog.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
