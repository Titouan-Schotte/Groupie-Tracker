package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func mapImage(lat, long float64) string {
	// Clé d'API Bing Maps
	apiKey := "98dHC1zw62rCO5MgbyLo~-RXs8b5NOfEDf1Ed_fpG5w~ApA31rcfZ3Il_YTnP5E7_VKZQYxqvk8eO5R2e5hzfQuR9jXpwfU_X5Y0wSv-K-iD"

	const bingMapsStaticURL = "https://dev.virtualearth.net/REST/v1/Imagery/Map/Road"
	return fmt.Sprintf("%s/%.6f,%.6f/16?mapSize=1000000,625000&pp=%.6f,%.6f;66&mapLayer=Basemap,Buildings&key=%s",
		bingMapsStaticURL, lat, long, lat, long, apiKey)
}

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
}

func getCoordinates(city string) (float64, float64) {
	// Call the API to get the coordinates of the city
	// Replace the API_URL with the appropriate URL for the Map API you are using
	apiURL := "https://api.api-ninjas.com/v1/geocoding?city=" + city
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiURL, nil)

	req.Header.Add("X-Api-Key", "VOdAuPTF1gLm8w2EIloGqw==w5vdhDxxJalElYDG")

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)

		// Unmarshal the JSON data into a slice of Location structs
		var locations []Location
		json.Unmarshal(bodyBytes, &locations)

		// Example of accessing the unmarshalled data
		for _, location := range locations {
			return location.Latitude, location.Longitude
		}
	} else {
		fmt.Println("Error:", resp.StatusCode)
	}
	return 0, 0
}

func GenerateMapImage(city string) fyne.Resource {
	// obtient les coordonnées de la ville
	latitude, longitude := getCoordinates(city)
	fmt.Println(latitude, longitude)
	imageURL := mapImage(latitude, longitude)
	// appel de la fonction pour obtenir l'image
	fmt.Println(imageURL)
	resource, _ := fyne.LoadResourceFromURLString(imageURL)

	// création de l'image
	image := canvas.NewImageFromResource(resource)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(200, 200)) // Ajuster la taille de l'image
	ressource, _ := fyne.LoadResourceFromURLString(imageURL)
	return ressource
}
