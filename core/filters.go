package core

import (
	"strconv"
	"strings"
)

// Filtrer par date de création
func FilterByCreationDate(artists []Artist, year int) []Artist {
	yearIn := int64(year)
	var filteredArtists []Artist
	for _, artist := range artists {
		if artist.CreationDate == yearIn {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// Filtrer par date du premier album
func FilterByFirstAlbumDate(artists []Artist, year int) []Artist {
	var filteredArtists []Artist

	for _, artist := range artists {
		intYear, _ := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])
		if intYear == year {
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
			if concert.Location.Locations[0] == location {
				filteredArtists = append(filteredArtists, artist)
				break // Une fois qu'un concert correspondant est trouvé, passer à l'artiste suivant
			}
		}
	}
	return filteredArtists
}
