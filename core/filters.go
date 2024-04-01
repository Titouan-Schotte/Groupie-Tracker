/*
Titouan Schotté
File with filters algorithm
*/
package core

import (
	"strconv"
	"strings"
)

// Filter by creation date
func FilterByCreationDate(artists []Artist, startYear, endYear int) []Artist {
	var filteredArtists []Artist

	for _, artist := range artists {
		// Extract the year of the first album
		albumYearStr := strconv.Itoa(int(artist.CreationDate))
		albumYear, _ := strconv.Atoi(albumYearStr)

		// Apply filtering logic with handling of -1 values
		if (startYear == -1 || albumYear >= startYear) && (endYear == -1 || albumYear <= endYear) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

func FilterByFirstAlbumDate(artists []Artist, startYear, endYear int) []Artist {
	var filteredArtists []Artist

	for _, artist := range artists {
		// Extract the year of the first album

		albumYearStr := strings.Split(artist.FirstAlbum, "-")[2]
		albumYear, _ := strconv.Atoi(albumYearStr)

		// Apply filtering logic with handling of -1 values
		if (startYear == -1 || albumYear >= startYear) && (endYear == -1 || albumYear <= endYear) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// Filter by number of members
func FilterByNumberOfMembers(artists []Artist, numMembers int) []Artist {
	var filteredArtists []Artist
	for _, artist := range artists {
		if len(artist.Members) == numMembers {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// Filter by concert locations
func FilterByConcertLocation(artists []Artist, location string) []Artist {
	var filteredArtists []Artist
	for _, artist := range artists {
		for _, concert := range artist.ConcertDates {
			if concert.Location.Locations[0] == location {
				filteredArtists = append(filteredArtists, artist)
				break // Once a matching concert is found, move on to the next artist
			}
		}
	}
	return filteredArtists
}
