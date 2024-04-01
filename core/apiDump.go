package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Artist represents an artist with their details.
type Artist struct {
	Id           int       `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Members      []string  `json:"members"`
	CreationDate int64     `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	Locations    string    `json:"locations"`
	ConcertDates []Concert `json:"concertDates"`
	Relations    string    `json:"relations"`
}

// Member represents a member of an artist group.
type Member struct {
	Surname string `json:"surname"`
	Name    string `json:"name"`
}

// Concert represents a concert with its date and location.
type Concert struct {
	Date     Date                `json:"dates"`
	Location APIResponseLocation `json:"locations"`
}

// APIResponseLocation represents the response structure for location data from the API.
type APIResponseLocation struct {
	Locations []string `json:"locations"`
}

// Date represents a date with day, month, and year.
type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// UnmarshalJSON is a custom unmarshal function for Date type to handle special cases.
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

// Api_artists retrieves artists' data from the API and shows it.
func Api_artists() []Artist {
	var response []Artist

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	defer res.Body.Close()

	body := newFunction(res)
	json.Unmarshal(body, &response)

	for i, p := range response {
		fmt.Printf("Artist %d: %s\n", i+1, p.Name)
		for j, concert := range p.ConcertDates {
			fmt.Printf("  Concert %d: Date - %d-%d-%d, Location - %v\n", j+1, concert.Date.Day, concert.Date.Month, concert.Date.Year, concert.Location.Locations)
		}
	}
	return response
}

// Api_dates retrieves dates data from the API and shows it.
func Api_dates() {
	var response4 map[string][]struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates")

	defer res.Body.Close()

	body := newFunction(res)
	json.Unmarshal(body, &response4)

	for i, item := range response4["index"] {
		for _, dateStr := range item.Dates {
			var date Date
			json.Unmarshal([]byte("\""+dateStr+"\""), &date)

			fmt.Printf("Date %d: ID: %d, Day: %d, Month: %d, Year: %d\n", i+1, item.Id, date.Day, date.Month, date.Year)
		}
	}
}

// Api_location retrieves location data from the API and shows it.
func Api_location() {
	var response2 APIResponseLocation

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	defer res.Body.Close()

	body := newFunction(res)
	json.Unmarshal(body, &response2)

	for i, location := range response2.Locations {
		fmt.Printf("Location %d: %s\n", i+1, location)
	}
}

// newFunction is a helper function to read response body from HTTP request.
func newFunction(res *http.Response) []byte {
	body, _ := ioutil.ReadAll(res.Body)

	return body
}
