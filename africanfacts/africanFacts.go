package africanfact

import (
	//impoer our encoding/package
	"container/list"
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

	//check if input in categorylist.
	// var intput string = "geography"

	// randomAfricanFact := giveRandomFactCategoryAfricanFactsByInputValue(intput, factCategory)
	// fmt.Println("facts:", randomAfricanFact.Fact)
	// fmt.Println("category:", randomAfricanFact.Category)

	//Give random african facts from the array.
	return giveRandomAfricanFacts(factCategory)

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

func giveRandomAfricanFacts(factCategory ArrayFactsCategory) FactCategory {

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

func giveRandomFactCategoryAfricanFactsByInputValue(inputValue string, factCategory ArrayFactsCategory) FactCategory {
	result := isValidCategory(inputValue)
	if result {
		fmt.Println(inputValue, "is present in the category")
		fmt.Println(".......................")
		//filter out the word in the arry of the category :input value.
		selectedAfricanFact := giveRandomFactCategoryAfricanFactsByCategory(inputValue, factCategory)
		fmt.Println("facts:", selectedAfricanFact.Fact)
		fmt.Println("category:", selectedAfricanFact.Category)
	}
	return factCategory.ArrayFactsCategory[0]
}

func giveRandomFactCategoryAfricanFactsByCategory(inputValue string, factCategory ArrayFactsCategory) FactCategory {

	selectedCategoryList := list.New()
	for i := 0; i < len(factCategory.ArrayFactsCategory); i++ {
		if inputValue == factCategory.ArrayFactsCategory[i].Category {
			selectedCategoryList.PushBack(factCategory.ArrayFactsCategory[i])
		}
	}

	min := 0
	max := selectedCategoryList.Len()
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	x := rand.Intn(max-min) + min

	selectedCategory := factCategory.ArrayFactsCategory[x]
	return selectedCategory
}
