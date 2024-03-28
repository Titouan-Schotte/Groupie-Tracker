package core

import (
	"strconv"
	"strings"
)

// Filtrer par date de création
func FilterByCreationDate(artists []Artist, startYear, endYear int) []Artist {
	var filteredArtists []Artist

	for _, artist := range artists {
		// Extraire l'année du premier album
		albumYearStr := strconv.Itoa(int(artist.CreationDate)) // Supposons que FirstAlbum soit au format DD-MM-YYYY
		albumYear, _ := strconv.Atoi(albumYearStr)
		// Appliquer la logique de filtrage avec gestion des valeurs -1
		if (startYear == -1 || albumYear >= startYear) && (endYear == -1 || albumYear <= endYear) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

func FilterByFirstAlbumDate(artists []Artist, startYear, endYear int) []Artist {
	var filteredArtists []Artist

	for _, artist := range artists {
		// Extraire l'année du premier album
		albumYearStr := strings.Split(artist.FirstAlbum, "-")[2] // Supposons que FirstAlbum soit au format DD-MM-YYYY
		albumYear, _ := strconv.Atoi(albumYearStr)
		// Appliquer la logique de filtrage avec gestion des valeurs -1
		if (startYear == -1 || albumYear >= startYear) && (endYear == -1 || albumYear <= endYear) {
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
		for _, concert := range artist.ConcertDates {
			if concert.Location == location {
				filteredArtists = append(filteredArtists, artist)
				break // Une fois qu'un concert correspondant est trouvé, passer à l'artiste suivant
			}
		}
	}
	return filteredArtists
}
