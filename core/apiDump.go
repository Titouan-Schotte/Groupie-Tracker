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
	Id           int       `json:"id"`
	Image        string    `json:"image"`
	Nom          string    `json:"name"`
	Members      []string  `json:"members"`
	CreationDate int64     `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	Locations    string    `json:"locations"`
	ConcertDates []Concert `json:"concertDates"`
	Relations    string    `json:"relations"`
}

type Member struct {
	Surname string `json:"surname"`
	Name    string `json:"name"`
}

type Concert struct {
	Date     Date   `json:"dates"`
	Location string `json:"locations"`
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

type RelationConcert struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func Api_artists() []Artist {
	var response []Artist

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	defer res.Body.Close()

	body := newFunction(res)
	json.Unmarshal(body, &response)

	for i, p := range response {
		var responseRel RelationConcert

		fmt.Printf("Artist %d: %s\n", i+1, p.Relations)
		resRel, _ := http.Get(p.Relations)

		defer resRel.Body.Close()

		bodyRel := newFunction(resRel) // Assurez-vous que cette fonction lit correctement le corps de la réponse
		json.Unmarshal(bodyRel, &responseRel)

		for location, dates := range responseRel.DatesLocations {
			for _, date := range dates { // Itérer sur chaque date dans la slice
				fmt.Printf("%s %s", location, date)
				dateSplit := strings.Split(date, "-")
				if len(dateSplit) == 3 { // Vérifiez que la date est bien formée
					dayIn, _ := strconv.Atoi(dateSplit[0])
					monthIn, _ := strconv.Atoi(dateSplit[1])
					yearIn, _ := strconv.Atoi(dateSplit[2])
					concertDate := Date{
						Day:   dayIn,
						Month: monthIn,
						Year:  yearIn,
					}
					response[i].ConcertDates = append(response[i].ConcertDates, Concert{Date: concertDate, Location: location})
				}
			}
		}
	}
	return response
}

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

func newFunction(res *http.Response) []byte {
	body, _ := ioutil.ReadAll(res.Body)

	return body
}
