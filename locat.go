package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Sites struct which contains
// an array of features
type Sites struct {
	Sites []Site `json:"features"`
}

// Site struct which contains a name
// a type and a list of geometry links
type Site struct {
	Type   string `json:"type"`
	Geometry Geometry `json:"geometry"`
	Param Param `json:"properties"`
}

// Geometry struct which contains a
// list of links
type Geometry struct {
	Type string `json:"type"`
	Coordinates []float32 `json:"coordinates"` 
}

// Param struct which contains a
// list of links
type Param struct {	
	Statel string `json:"location_state"`
	City string `json:"city"`
	Zipcode string `json:"zipcode"`
	Loczip string `json:"location_zip"`
	Cupsid string `json:"cleanup_site_id"`
	Latitude string `json:"latitude"`
	County string `json:"county_name"`
	Locadd string `json:"location_address"`
	Cupname string `json:"cleanup_site_name"`
	Loccity string `json:"location_city"`
	Stat string `json:"statute"`
	Longitude string `json:"longitude"`
	Cupstatus string `json:"cleanup_status"`
	Address string `json:"address"`
	Facid string `json:"facility_site_id"`	
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("response.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened response.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Sites array
	var features Sites

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'features' which we defined above
	json.Unmarshal(byteValue, &features)
	
	// we iterate through every user within our features array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(features.Sites); i++ {
	  	fmt.Println("\n")	
		fmt.Println("Type: " + features.Sites[i].Type)
		fmt.Println("City: " + features.Sites[i].Param.City)
		fmt.Println("Zipcode: " + features.Sites[i].Param.Zipcode)
		fmt.Println("Longitude: " + features.Sites[i].Param.Longitude)
		fmt.Println("Latitude: " + features.Sites[i].Param.Latitude)
		fmt.Println("Status: " + features.Sites[i].Param.Cupstatus)
		fmt.Println("Site ID: " + features.Sites[i].Param.Cupsid)
		fmt.Println("Facility ID: " + features.Sites[i].Param.Facid)
		fmt.Println("Site Name: " + features.Sites[i].Param.Cupname)
		fmt.Println("Address: " + features.Sites[i].Param.Address)		
	  	fmt.Printf("lon: %f \n",features.Sites[i].Geometry.Coordinates[0])
	  	fmt.Printf("lat: %f \n",features.Sites[i].Geometry.Coordinates[1])
	}
	  	fmt.Println("\n")	
}

