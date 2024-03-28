package core

//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//
//type GeoResponse struct {
//	Latitude  float64 `json:"latitude"`
//	Longitude float64 `json:"longitude"`
//}
//
//func MapAPI() (float64, float64) {
//	address := "Mainz, Germany"
//	apiKey := "YOUR_API_KEY" // Remplacez YOUR_API_KEY par votre clé d'API
//
//	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", address, apiKey)
//
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println("Erreur lors de la requête HTTP:", err)
//		return 0, 0
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Erreur lors de la lecture de la réponse:", err)
//		return 0, 0
//	}
//
//	var geoResp GeoResponse
//	err = json.Unmarshal(body, &geoResp)
//	if err != nil {
//		fmt.Println("Erreur lors de la conversion JSON:", err)
//		return 0, 0
//	}
//
//	return geoResp.Latitude, geoResp.Longitude
//}

import (
	"fmt"
)

func main() {
	// Clé d'API Bing Maps
	apiKey := "98dHC1zw62rCO5MgbyLo~-RXs8b5NOfEDf1Ed_fpG5w~ApA31rcfZ3Il_YTnP5E7_VKZQYxqvk8eO5R2e5hzfQuR9jXpwfU_X5Y0wSv-K-iD"

	// Paramètres de la requête
	Latitude := 51.5074
	Longitude := -0.1278

	const bingMapsStaticURL = "https://dev.virtualearth.net/REST/v1/Imagery/Map/Road"
	fmt.Printf("%s/%.6f,%.6f/16?mapSize=800,500&pp=%.6f,%.6f;66&mapLayer=Basemap,Buildings&key=%s",
		bingMapsStaticURL, Latitude, Longitude, Latitude, Longitude, apiKey)
}
