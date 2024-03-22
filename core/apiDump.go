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
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Nom          string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates []Concert
	Relations    string `json:"relations"`
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

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error getting artists:", err)
		return nil
	}
	defer res.Body.Close()

	body := newFunction(res)

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling artist response:", err)
		return nil
	}

	for i, p := range response {
		fmt.Printf("Artist %d: %s\n", i+1, p.Nom)
		for j, concert := range p.ConcertDates {
			fmt.Printf("  Concert %d: Date - %d-%d-%d, Location - %v\n", j+1, concert.Date.Day, concert.Date.Month, concert.Date.Year, concert.Location.Locations)
		}
	}
	return response
}

func Api_dates() {
	var response4 map[string][]struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("Error getting dates:", err)
		return
	}
	defer res.Body.Close()

	body := newFunction(res)

	if err := json.Unmarshal(body, &response4); err != nil {
		fmt.Println("Error unmarshalling dates response:", err)
		return
	}

	for i, item := range response4["index"] {
		for _, dateStr := range item.Dates {
			var date Date
			if err := json.Unmarshal([]byte("\""+dateStr+"\""), &date); err != nil {
				fmt.Println("Error unmarshalling date:", err)
				continue
			}
			fmt.Printf("Date %d: ID: %d, Day: %d, Month: %d, Year: %d\n", i+1, item.Id, date.Day, date.Month, date.Year)
		}
	}
}

func Api_location() {
	var response2 APIResponseLocation

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("Error getting locations:", err)
		return
	}
	defer res.Body.Close()

	body := newFunction(res)

	if err := json.Unmarshal(body, &response2); err != nil {
		fmt.Println("Error unmarshalling locations response:", err)
		return
	}

	for i, location := range response2.Locations {
		fmt.Printf("Location %d: %s\n", i+1, location)
	}
}

func newFunction(res *http.Response) []byte {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	return body
}
