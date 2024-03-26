package commands

import (
	"errors"
	"fmt"
	"github.com/Taliker/Gokedex/internal/api"
	"time"
)

func CommandExplore(config *Config, areaName string) error {
	if areaName == "" {
		return errors.New("please enter the name of the area you want to explore")
	}

	location, err := api.GetLocation(areaName, &config.Cache)
	if err != nil {
		return err
	}
	printLocation(location)
	return nil
}

func printLocation(location api.LocationArea) {
	fmt.Printf("Exploring %s ...", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range location.Encounters {
		fmt.Println(pokemon.Name)
		time.Sleep(100 * time.Millisecond)
	}
}
