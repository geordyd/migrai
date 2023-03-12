package main

import (
	//impoer our encoding/package
	"encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "log"
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
    // check which value of category 

   var factCategory ArrayFactsCategory 
   var catergoryList = [6] string {"geography","wildlife","demographics","language","economy","nature"}
   fmt.Println(catergoryList)
	
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
 
    // Let's print the unmarshalled data!
    for i := 0; i < len(factCategory.ArrayFactsCategory); i++ {
        fmt.Println("category: " + factCategory.ArrayFactsCategory[i].Fact)
        fmt.Println("facts: " + factCategory.ArrayFactsCategory[i].Category)
    }

}