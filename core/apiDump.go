package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Artist struct {
	Id           int
	Image        string
	Nom          string
	Members      []string
	CreationDate int64
	FirstAlbum   string
	Concerts     []Concert
	Relations    string
}
type Member struct {
	Surname string `json:"surname"`
	Name    string `json:"name"`
}

type Concert struct {
	Date     Date                `json:"dates"`
	Location APIResponseLocation `json:"locations"`
}

type APIResponseLocation struct {
	Locations []string `json:"locations"`
}

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type API struct {
}

type APIResponseDates struct {
	Dates []string `json:"dates"`
}

func (d *Date) UnmarshalJSON(data []byte) {
	var dateStr string
	json.Unmarshal(data, &dateStr)
	hasAsterisk := strings.HasPrefix(dateStr, "*")

	if hasAsterisk {
		dateStr = dateStr[1:]
	}

	dateComponents := strings.Split(dateStr, "-")

	if len(dateComponents) == 3 {
		d.Day, _ = strconv.Atoi(dateComponents[0])
		d.Month, _ = strconv.Atoi(dateComponents[1])
		d.Year, _ = strconv.Atoi(dateComponents[2])
	}

	return
}

func Api_artists() []Artist {
	var response []Artist

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	defer res.Body.Close()

	body := newFunction(res)

	json.Unmarshal(body, &response)

	for i, p := range response {
		fmt.Printf("test %d: %s, %v, %d, %s, %v, %v\n", i+1, p.Nom, p.Members, p.CreationDate, p.FirstAlbum, p.Concerts, p.Image)

		if len(p.Concerts) == 0 {
			fmt.Println("No concerts available")
		} else {
			for j, concert := range p.Concerts {
				fmt.Printf("  Concert %d: Date - %d-%d-%d, Location - %v\n", j+1, concert.Date.Day, concert.Date.Month, concert.Date.Year, concert.Location.Locations)
			}
		}

		fmt.Printf("%v\n", p.Image)
	}
	return response
}

func Api_dates() {
	var response4 APIResponseDates

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	defer res.Body.Close()

	body := newFunction(res)

	json.Unmarshal(body, &response4)
	var rawResponse map[string][]struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}

	json.Unmarshal(body, &rawResponse)

	for i, item := range rawResponse["index"] {
		for _, dateStr := range item.Dates {
			var date Date
			json.Unmarshal([]byte("\""+dateStr+"\""), &date)
			fmt.Printf("test %d: ID: %d, Day: %d, Month: %d, Year: %d\n", i+1, item.Id, date.Day, date.Month, date.Year)
		}
	}
}

func Api_location() {
	var response2 APIResponseLocation

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	defer res.Body.Close()

	body := newFunction(res)

	json.Unmarshal(body, &response2)

	for i, location := range response2.Locations {
		fmt.Printf("test %d: %s\n", i+1, location)
	}
}

func newFunction(res *http.Response) []byte {
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
