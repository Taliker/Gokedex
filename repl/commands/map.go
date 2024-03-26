package commands

import (
	"errors"
	"fmt"
	"github.com/Taliker/Gokedex/internal/api"
	"time"
)

func CommandMapF(config *Config, arg string) error {
	if config.nextURL == "" {
		return commandMap("", config)
	} else {
		return commandMap(config.nextURL, config)
	}
}

func CommandMapB(config *Config, arg string) error {
	if config.prevURL == "" {
		return errors.New("no previous locations, please use mapf to get locations")
	} else {
		return commandMap(config.prevURL, config)
	}
}

func commandMap(url string, config *Config) error {
	if locations, err := api.GetLocations(url, &config.Cache); err == nil {
		config.nextURL = locations.Next
		config.prevURL = locations.Previous
		printLocations(locations.Results)
	}
	return nil
}

func printLocations(locations []api.Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
		time.Sleep(100 * time.Millisecond)
	}
}
