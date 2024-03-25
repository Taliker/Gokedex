package commands

import (
	"errors"
	"fmt"
	locPac "github.com/Taliker/Gokedex/internal/http/locations"
)

const (
	locationURL = "https://pokeapi.co/api/v2/location/"
)

func CommandMapF(config *Config) error {
	if config.nextURL == "" {
		return commandMap(locationURL, config)
	} else {
		return commandMap(config.nextURL, config)
	}
}

func CommandMapB(config *Config) error {
	if config.prevURL == "" {
		return errors.New("no previous locations, please use mapf first to get locations")
	} else {
		return commandMap(config.prevURL, config)
	}
}

func commandMap(url string, config *Config) error {
	locations := locPac.GetLocations(url)
	config.nextURL = locations.Next
	config.prevURL = locations.Previous
	printLocations(locations.Results)
	return nil
}

func printLocations(locations []locPac.Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
