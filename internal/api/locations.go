package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Taliker/Gokedex/internal/cache"
	"io"
	"net/http"
)

const (
	locationURL = "https://pokeapi.co/api/v2/location-area/"
)

func GetLocation(areaName string, cache *cache.Cache) (LocationArea, error) {
	if areaName == "" {
		return LocationArea{}, errors.New("please enter the name of the area you want to explore")
	}
	url := locationURL + areaName
	if val, ok := cache.Get(url); ok {
		//From cache
		fmt.Println("Cache")
		var location LocationArea
		location, err := LocationAreaDataToLocationArea(val)
		if err != nil {
			return LocationArea{}, errors.New("failed to unmarshal cached response")
		}
		return location, nil
	} else {
		//From API
		res, err := http.Get(url)
		if err != nil {
			return LocationArea{}, err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return LocationArea{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
		}
		if err != nil {
			return LocationArea{}, errors.New("failed to read response body")
		}

		cache.Add(url, body)

		var location LocationArea
		location, err = LocationAreaDataToLocationArea(body)
		if err != nil {
			return LocationArea{}, errors.New("failed to unmarshal response body")
		}
		return location, nil
	}
}

func GetLocations(url string, cache *cache.Cache) (LocationsResult, error) {
	usedURL := url
	if url == "" {
		usedURL = locationURL + "?limit=10"
	}
	if val, ok := cache.Get(usedURL); ok {
		//From cache
		var locations LocationsResult
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return LocationsResult{}, errors.New("failed to unmarshal cached response")
		}
		return locations, nil
	} else {
		//From API
		res, err := http.Get(usedURL)
		if err != nil {
			return LocationsResult{}, err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return LocationsResult{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
		}
		if err != nil {
			return LocationsResult{}, errors.New("failed to read response body")
		}

		cache.Add(usedURL, body)

		var locations LocationsResult
		err = json.Unmarshal(body, &locations)
		if err != nil {
			return LocationsResult{}, errors.New("failed to unmarshal response body")
		}
		return locations, nil
	}
}

type LocationsResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Location
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationArea struct {
	Name       string
	Encounters []Pokemon
}
