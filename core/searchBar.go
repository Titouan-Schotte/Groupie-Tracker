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

		// Calcul du potentiel
		potIn := 0
		for _, el := range artistQuery {
			if strings.Contains(el, inp) || strings.Contains(inp, el) {
				potIn++
			}
		}
		potentials = append(potentials, PotentialSearch{Potential: potIn, Query: artistQuery})
	}

	// Trie des résultats potentiels par ordre décroissant de potentiel
	sort.Slice(potentials, func(i, j int) bool {
		return potentials[i].Potential > potentials[j].Potential
	})

	// Récupération des artistes correspondant aux résultats les plus pertinents
	var resultArtists []Artist
	for _, potential := range potentials {
		if potential.Potential > 0 {
			// Parcours de tous les artistes pour vérifier s'ils correspondent au résultat potentiel
			for _, artist := range artists {
				// Vérification si l'artiste correspond au résultat potentiel
				if containsAll(artist, potential.Query) {
					// Si l'artiste correspond, l'ajouter à la slice resultArtists
					resultArtists = append(resultArtists, artist)
					// Break pour passer au résultat potentiel suivant
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
