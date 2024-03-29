/*
Titouan Schotté
Search bar algorithms
*/
package core

import (
	"sort"
	"strconv"
	"strings"
)

type PotentialSearch struct {
	Query     []string
	Potential int
}

func SearchInAllStruct(inp string, artists []Artist) []Artist {
	inp = strings.ToLower(inp)
	potentials := []PotentialSearch{}
	for _, artist := range artists {
		convCreationDate := strconv.Itoa(int(artist.CreationDate))
		artistQuery := []string{strings.ToLower(artist.Nom), convCreationDate}
		for _, member := range artist.Members {
			artistQuery = append(artistQuery, strings.ToLower(member))
		}
		for _, concert := range artist.ConcertDates {
			artistQuery = append(artistQuery, strconv.Itoa(concert.Date.Day))
			artistQuery = append(artistQuery, strconv.Itoa(concert.Date.Month))
			artistQuery = append(artistQuery, strconv.Itoa(concert.Date.Year))
			artistQuery = append(artistQuery, strings.ToLower(concert.Location))
		}

		//Here we operate on the principle that the closer an artist is to the input, the more potential he has.
		//The potential is an integer data.
		// Calcul du potentiel
		potIn := 0
		for _, el := range artistQuery {
			if strings.Contains(el, inp) || strings.Contains(inp, el) {
				potIn++
			}
		}
		potentials = append(potentials, PotentialSearch{Potential: potIn, Query: artistQuery})
	}

	// Sorts potential results in descending order of potential
	sort.Slice(potentials, func(i, j int) bool {
		return potentials[i].Potential > potentials[j].Potential
	})

	// Retrieval of artists corresponding to the most relevant results
	var resultArtists []Artist
	for _, potential := range potentials {
		if potential.Potential > 0 {

			// Run through all artists to check if they match the potential outcome
			for _, artist := range artists {
				// Checking if the artist matches the potential result
				if containsAll(artist, potential.Query) {
					// If the artist matches, add it to the resultArtists slice
					resultArtists = append(resultArtists, artist)
					break
				}
			}
		}
	}

	return resultArtists
}

func containsAll(artist Artist, query []string) bool {
	for _, q := range query {
		found := false
		for _, field := range artistData(artist) {
			if strings.Contains(field, q) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func artistData(artist Artist) []string {
	convCreationDate := strconv.Itoa(int(artist.CreationDate))
	data := []string{strings.ToLower(artist.Nom), convCreationDate}
	for _, member := range artist.Members {
		data = append(data, strings.ToLower(member))
	}
	for _, concert := range artist.ConcertDates {
		data = append(data, strconv.Itoa(concert.Date.Day), strconv.Itoa(concert.Date.Month), strconv.Itoa(concert.Date.Year), strings.ToLower(concert.Location))
	}
	return data
}
