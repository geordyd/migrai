import (
	//impoer our encoding/package
	"encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)


// Description a struct whih contains category and facts.
type FactCategory struct {
	Category string `json:"category"`
	Fact string `json:"fact"`
}


func main() {

    // Open our jsonFile
    jsonFile, err := os.Open("facts.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    fmt.Println(result["facts"])

}