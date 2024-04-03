package core

import (
	"encoding/json"
	"os"
)

type Favoris struct {
	Favoris []struct {
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
}

var Fav Favoris

func AddFavorite(artist Artist) { // Add an artist to the favorite list
	for i := 0; i < len(Fav.Favoris); i++ {
		if Fav.Favoris[i].Id == artist.Id {
			return
		}
	}
	Fav.Favoris = append(Fav.Favoris, artist)
	Updatejson()
}

func RemoveFavorite(artist Artist) { // Remove an artist from the favorite list
	for i, a := range Fav.Favoris {
		if a.Nom == artist.Nom {
			Fav.Favoris = append(Fav.Favoris[:i], Fav.Favoris[i+1:]...)
			break
		}
	}
	Updatejson()
}

func InFavorite(artist Artist) bool { // Check if an artist is in the favorite list
	for i := 0; i < len(Fav.Favoris); i++ {
		if Fav.Favoris[i].Id == artist.Id {
			return true
		}
	}
	return false
}

func GetFavorite() []Artist { // Get the favorite list
	var artists []Artist
	for i := 0; i < len(Fav.Favoris); i++ {
		artists = append(artists, Fav.Favoris[i])
	}
	return artists
}

func Updatejson() { // Update the json file
	data, _ := json.Marshal(Fav)
	os.WriteFile("./DataBase/Favoris.json", data, 0644)
}

func LoadJson() { // Load the json file
	file, _ := os.ReadFile("./DataBase/Favoris.json")
	json.Unmarshal(file, &Fav)
}
