package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	hasAsterisk := strings.HasPrefix(dateStr, "*")

	if hasAsterisk {
		dateStr = dateStr[1:]
	}

	dateComponents := strings.Split(dateStr, "-")

	if len(dateComponents) == 3 {
		var err error
		d.Day, err = strconv.Atoi(dateComponents[0])
		if err != nil {
			return err
		}

		d.Month, err = strconv.Atoi(dateComponents[1])
		if err != nil {
			return err
		}

		d.Year, err = strconv.Atoi(dateComponents[2])
		if err != nil {
			return err
		}
	}

	return nil
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
	var rawResponse map[string][]struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}

	err = json.Unmarshal(body, &rawResponse)
	if err != nil {
		log.Fatal(err)
	}

	for i, item := range rawResponse["index"] {
		for _, dateStr := range item.Dates {
			var date Date
			err := json.Unmarshal([]byte("\""+dateStr+"\""), &date)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("test %d: ID: %d, Day: %d, Month: %d, Year: %d\n", i+1, item.Id, date.Day, date.Month, date.Year)
		}
	}
}

func Api_location() {
	var response2 APIResponseLocation

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
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
		log.Fatal(err)
	}

	for i, location := range response2.Locations {
		fmt.Printf("test %d: %s\n", i+1, location)
	}
}

func newFunction(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
