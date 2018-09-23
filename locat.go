package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)


// Sites struct which contains
// an array of cupsites
type Sites struct {
	Sites []Site `json:"cupsites"`
}


// Site struct which contains a name
// a type and a list of social links
type Site struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
	Properties string `json:"properties"`
	Param Param `json:"params"`
}


// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter string `json:"twitter"`
}

// Param struct which contains a
// list of links
type Param struct {
	Street string `json:"street"`
	Color string `json:"color"`
}

func main() {


	// Open our jsonFile
	jsonFile, err := os.Open("response.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}


	fmt.Println("Successfully Opened cupsites.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()


	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)


	// we initialize our Sites array
	var cupsites Sites


	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'cupsites' which we defined above
	json.Unmarshal(byteValue, &cupsites)


	// we iterate through every user within our cupsites array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(cupsites.Sites); i++ {
		fmt.Println("Site Type: " + cupsites.Sites[i].Type)
		fmt.Println("Site Age: " + strconv.Itoa(cupsites.Sites[i].Age))
		fmt.Println("Site Name: " + cupsites.Sites[i].Name)
		fmt.Println("Facebook Url: " + cupsites.Sites[i].Social.Facebook)
		fmt.Println("City: " + cupsites.Sites[i].Properties)		
		fmt.Println("Color: " + cupsites.Sites[i].Param.Color)
	}


}

