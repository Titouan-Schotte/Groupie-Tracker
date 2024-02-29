package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type APIResponse struct {
	Artists   []Artist    `json:"artists"`
	Locations []Locations `json:"locations"`
	Dates     []Dates     `json:"dates"`
	Relation  []Relation  `json:"relation"`
}

type APIResponseDates struct {
	Dates []string `json:"dates"`
}

type APIResponseLocation struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Artist struct {
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int64       `json:"id"`
	DatesLocations []Locations `json:"datesLocations"`
}

func Api_artists() {
	var response []Artist

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := newFunction(res)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	for i, p := range response {
		fmt.Printf("test %d: %s, %s, %d, %s, %s, %s\n", i+1, p.Name, p.Members, p.CreationDate, p.FirstAlbum, p.ConcertDates, p.Image)
	}
}

func Api_dates() {
	var response4 APIResponseDates

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := newFunction(res)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response4)
	if err != nil {
		log.Fatal(err)
	}

	for i, date := range response4.Dates {
		fmt.Printf("test %d: %s\n", i+1, date)
	}
}
func Api_location() {
	var response3 []APIResponseLocation

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := newFunction(res)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response3)
	if err != nil {
		var singleResponse APIResponseLocation
		err = json.Unmarshal(body, &singleResponse)
		if err != nil {
			log.Fatal(err)
		}
		response3 = append(response3, singleResponse)
	}

	for i, loc := range response3 {
		fmt.Printf("test %d: %s, %s\n", i+1, loc.Locations, loc.Dates)
	}
}

func Api_Relation() {
	var response2 []Relation

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := newFunction(res)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response2)
	if err != nil {
		var singleResponse Relation
		err = json.Unmarshal(body, &singleResponse)
		if err != nil {
			log.Fatal(err)
		}
		response2 = append(response2, singleResponse)
	}

	for i, p := range response2 {
		fmt.Printf("test %d:%v\n", i+1, p.DatesLocations)
	}
}

func newFunction(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
