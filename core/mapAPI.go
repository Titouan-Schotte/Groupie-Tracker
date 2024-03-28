package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func mapImage(lat, long float64) string {
	// Clé d'API Bing Maps
	apiKey := "98dHC1zw62rCO5MgbyLo~-RXs8b5NOfEDf1Ed_fpG5w~ApA31rcfZ3Il_YTnP5E7_VKZQYxqvk8eO5R2e5hzfQuR9jXpwfU_X5Y0wSv-K-iD"

	const bingMapsStaticURL = "https://dev.virtualearth.net/REST/v1/Imagery/Map/Road"
	return fmt.Sprintf("%s/%.6f,%.6f/16?mapSize=800,500&pp=%.6f,%.6f;66&mapLayer=Basemap,Buildings&key=%s",
		bingMapsStaticURL, lat, long, lat, long, apiKey)
}

func getCoordinates(city string) (float64, float64, error) {
	// Call the API to get the coordinates of the city
	// Replace the API_URL with the appropriate URL for the Map API you are using
	API_URL := "https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1" + city
	response, err := http.Get(API_URL)
	if err != nil {
		return 0, 0, err
	}
	defer response.Body.Close()

	// Parse the response to extract the latitude and longitude
	var data struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return 0, 0, err
	}

	return data.Latitude, data.Longitude, nil
}

func GenerateMapImage(city string) (fyne.Resource, error) {
	// obtient les coordonnées de la ville
	latitude, longitude, err := getCoordinates(city)
	if err != nil {
		return nil, err
	}
	imageURL := mapImage(latitude, longitude)
	// appel de la fonction pour obtenir l'image
	resource, err := fyne.LoadResourceFromURLString(imageURL)
	if err != nil {
		return nil, err
	}

	// création de l'image
	image := canvas.NewImageFromResource(resource)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(200, 200)) // Ajuster la taille de l'image

	return fyne.LoadResourceFromURLString(imageURL)
}
