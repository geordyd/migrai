package africanfact

import (
	//impoer our encoding/package

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type ArrayFactsCategory struct {
	ArrayFactsCategory []FactCategory `json:"data"`
}

// Description a struct with contains category and facts.
type FactCategory struct {
	Category string `json:"category"`
	Fact     string `json:"fact"`
}

func Africanfact() FactCategory {
	// get the inpute value of the user and set in a variable.
	// check if input value contains in category list
	var factCategory ArrayFactsCategory
	var catergoryList = []string{"geography", "wildlife", "demographics", "language", "economy", "nature"}
	fmt.Println(catergoryList)
	fmt.Println("Pick out which category you want from the category list")
	fmt.Println("Type category which african facts you want to see with '!african-facts 'given category' command.\n")
	fmt.Println("```Available categories:\n - 1.geography\n - 2.wildlife\n - 3.demographics\n - 4.language\n - 5.economy\n - 6.nature\n```")
	fmt.Println("Give a random pick from african facts with '!african-facts random' command.\n")
	readOutFileInArray(factCategory)
	//Give random african facts from the array.
	return giveRandomAfricanFacts(factCategory)
}

func readOutFileInArray(factCategory ArrayFactsCategory) {
	// Open our jsonFile
	jsonFile, err := os.Open("facts.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened facts.json!")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	content, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &factCategory)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}

func isValidCategory(category string) bool {
	switch category {
	case
		"wildlife",
		"demographics",
		"language",
		"economy",
		"nature",
		"geography":
		return true
	}
	return false
}

func randomNizer(factCategory ArrayFactsCategory) FactCategory {
	//give random facts of the african-fact list
	min := 0
	max := len(factCategory.ArrayFactsCategory)

	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	x := rand.Intn(max-min) + min

	//show the random pick from africanfacts list.
	fmt.Println(".......................")
	fmt.Println("facts:", factCategory.ArrayFactsCategory[x].Fact)
	fmt.Println("category:", factCategory.ArrayFactsCategory[x].Category)
	return factCategory.ArrayFactsCategory[x]
}

func giveRandomAfricanFacts(factCategory ArrayFactsCategory) FactCategory {
	return randomNizer(factCategory)
}
