package mugshot

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type MugshotData struct {
	ID          string
	Sex         string
	Height      string
	Weight      string
	Hair        string
	Eyes        string
	Race        string
	SexOffender string
	Offense     string
}

func Get() ([]io.Reader, MugshotData, error) {

	csv, err := readCsv()
	if err != nil {
		fmt.Println(err)
		return nil, MugshotData{}, err
	}

	index := randomIndex(csv)

	randomMugshot := csv[index]

	mugshot := MugshotData{
		ID:          randomMugshot[0],
		Sex:         randomMugshot[1],
		Height:      randomMugshot[2],
		Weight:      randomMugshot[3],
		Hair:        randomMugshot[4],
		Eyes:        randomMugshot[5],
		Race:        randomMugshot[6],
		SexOffender: randomMugshot[7],
		Offense:     randomMugshot[8],
	}

	images := getImages(mugshot.ID)

	return images, mugshot, nil
}

func randomIndex(csv [][]string) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	min := 1
	max := len(csv) - 1

	randomIndex := r1.Intn(max-min) + min

	return randomIndex
}

func readCsv() ([][]string, error) {
	csvFile, err := os.Open("mugshots/labels_utf8.csv")
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	return csvLines, nil
}

func getImages(imageName string) []io.Reader {
	var images []io.Reader

	image1, err := os.Open(fmt.Sprintf("mugshots/front/%s", imageName))
	if err != nil {
		fmt.Println("Image1 doesn't exist")
	} else {
		images = append(images, image1)
	}

	image2, err := os.Open(fmt.Sprintf("mugshots/side/%s", imageName))
	if err != nil {
		fmt.Println("Image2 doesn't exist")
	} else {
		images = append(images, image2)
	}
	return images
}
