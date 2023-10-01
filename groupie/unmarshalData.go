package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func readThefile(url string) []byte {
	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Print raw JSON response
	return body
	// Parse JSON
}

func UnmarshalData() ([]Artist, ResponseData, ResponseData2, ResponseData4, bool) {
	flag := true
	//for artist
	url := "https://groupietrackers.herokuapp.com/api/artists"
	body := readThefile(url)
	var artists []Artist
	err := json.Unmarshal(body, &artists)
	if err != nil {
		flag = false
		log.Fatal(err)
	}

	//for the locations
	url = "https://groupietrackers.herokuapp.com/api/locations"
	data := readThefile(url)
	var TheLocations ResponseData
	err = json.Unmarshal(data, &TheLocations)
	if err != nil {
		flag = false
		log.Fatal(err)
	}

	for z := 0; z < len(TheLocations.Index); z++ {
		for i := 0; i < len(TheLocations.Index[z].TheData); i++ {
			TheLocations.Index[z].TheData[i] = strings.ReplaceAll(TheLocations.Index[z].TheData[i], "_", " ")
		}
	}

	// for the dates
	url = "https://groupietrackers.herokuapp.com/api/dates"
	Datees := readThefile(url)
	var TheDates ResponseData2
	err = json.Unmarshal(Datees, &TheDates)
	if err != nil {
		flag = false
		log.Fatal(err)
	}
	// for the relations
	url = "https://groupietrackers.herokuapp.com/api/relation"
	relations := readThefile(url)
	var TheRelations ResponseData4
	err = json.Unmarshal(relations, &TheRelations)
	if err != nil {
		log.Fatal(err)
	}

	return artists, TheLocations, TheDates, TheRelations, flag
}
