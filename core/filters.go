package core

import (
	"strconv"
	"strings"
)

// Filtrer par date de création
func FilterByCreationDate(artists []Artist, year int) []Artist {
	var filteredArtists []Artist
	for _, artist := range artists {
		creationYear, _ := strconv.Atoi(strings.Split(artist.CreationDate, "/")[2])
		if creationYear == year {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// Filtrer par date du premier album
func FilterByFirstAlbumDate(artists []Artist, year int) []Artist {
	var filteredArtists []Artist
	for _, artist := range artists {
		if artist.FirstAlbum.Year == year {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// Filtrer par nombre de membres
func FilterByNumberOfMembers(artists []Artist, numMembers int) []Artist {
	var filteredArtists []Artist
	for _, artist := range artists {
		if len(artist.Members) == numMembers {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// Filtrer par lieux des concerts
func FilterByConcertLocation(artists []Artist, location string) []Artist {
	var filteredArtists []Artist
	for _, artist := range artists {
		for _, concert := range artist.Concerts {
			if concert.Location == location {
				filteredArtists = append(filteredArtists, artist)
				break // Une fois qu'un concert correspondant est trouvé, passer à l'artiste suivant
			}
		}
	}
	return filteredArtists
}
