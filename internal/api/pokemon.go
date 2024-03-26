package api

import (
	"errors"
	"fmt"
	"github.com/Taliker/Gokedex/internal/cache"
	"io"
	"net/http"
)

const (
	pokemonURL = "https://pokeapi.co/api/v2/pokemon/"
)

func GetPokemon(pokemonName string, cache *cache.Cache) (Pokemon, error) {
	if pokemonName == "" {
		return Pokemon{}, errors.New("please enter the name of the area you want to explore")
	}
	url := pokemonURL + pokemonName
	if val, ok := cache.Get(url); ok {
		//From cache
		fmt.Println("Cache")
		var pokemon Pokemon
		pokemon, err := PokemonDataFromJSON(val)
		if err != nil {
			return Pokemon{}, errors.New("failed to unmarshal cached response")
		}
		return pokemon, nil
	} else {
		//From API
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
		}
		if err != nil {
			return Pokemon{}, errors.New("failed to read response body")
		}

		cache.Add(url, body)

		var pokemon Pokemon
		pokemon, err = PokemonDataFromJSON(body)
		if err != nil {
			return Pokemon{}, errors.New("failed to unmarshal response body")
		}
		return pokemon, nil
	}
}
