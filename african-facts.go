package main

import (
	//impoer our encoding/package
	"encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "log"
    "math/rand"
    "time"
);

type ArrayFactsCategory struct {
    ArrayFactsCategory []FactCategory `json:"facts"`
}

// Description a struct whih contains category and facts.
type FactCategory struct {
	Category string `json:"category"`
	Fact string `json:"fact"`
}


func main() {
	
    // get the inpute value of the user and set in a variable.
    // check if input value contains in category list 
   var result bool = false
   var factCategory ArrayFactsCategory 
   var catergoryList = [6] string {"geography","wildlife","demographics","language","economy","nature"}
   fmt.Println(catergoryList)
   fmt.Println("Pick out which category you want from the category list")
   fmt.Println( "```Available categories:\n - 1.geography\n - 2.wildlife\n - 3.demographics\n - 4.language\n - 5.economy\n - 6.nature\n```")
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
    var intput string = "geography";
    result = IsValidCategory(intput)
    if result {
        fmt.Println(intput, "is present in the array of strings", catergoryList)
      

    }
    
    //Give random african facts from the array.
    giveRandomAfricanFacts(factCategory)
     
}

func IsValidCategory(category string) bool {
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

func giveRandomAfricanFacts(factCategory ArrayFactsCategory) {
    
    //give random facts of the african-fact list
    min := 0
    max := len(factCategory.ArrayFactsCategory)

    // set seed
    rand.Seed(time.Now().UnixNano())
    // generate random number and print on console
    x := rand.Intn(max - min) + min

    //show the random pick from africanfacts list. 
    fmt.Println("facts:",factCategory.ArrayFactsCategory[x].Fact)
    fmt.Println("category:",factCategory.ArrayFactsCategory[x].Category)
}

func giveRandomAfricanFactsByInput( inputValue string ,factCategory ArrayFactsCategory) FactCategory {
   //filter out the word in the arry of the category :input value.

   
     
}