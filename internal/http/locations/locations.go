package locations

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Taliker/Gokedex/internal/cache"
)

func GetLocations(url string, cache *cache.Cache) (LocationResult, error) {
	if val, ok := cache.Get(url); ok {
		//From cache
		var locations LocationResult
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return LocationResult{}, errors.New("failed to unmarshal cached response")
		}
		return locations, nil
	} else {
		//From API
		time.Sleep(2000 * time.Millisecond)
		res, err := http.Get(url)
		if err != nil {
			return LocationResult{}, err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return LocationResult{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
		}
		if err != nil {
			return LocationResult{}, errors.New("failed to read response body")
		}

		cache.Add(url, body)

		var locations LocationResult
		err = json.Unmarshal(body, &locations)
		if err != nil {
			return LocationResult{}, errors.New("failed to unmarshal response body")
		}
		return locations, nil
	}
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
