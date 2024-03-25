package locations

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetLocations(url string) LocationResult {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locations LocationResult
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}
	return locations
}

type LocationResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Location
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
