package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GeoResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func MapAPI() (float64, float64) {
	address := "Mainz, Germany"
	apiKey := "YOUR_API_KEY" // Remplacez YOUR_API_KEY par votre clé d'API

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", address, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP:", err)
		return 0, 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return 0, 0
	}

	var geoResp GeoResponse
	err = json.Unmarshal(body, &geoResp)
	if err != nil {
		fmt.Println("Erreur lors de la conversion JSON:", err)
		return 0, 0
	}

	return geoResp.Latitude, geoResp.Longitude
}
