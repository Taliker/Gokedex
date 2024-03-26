package commands

import (
	"errors"
	"fmt"
	"github.com/Taliker/Gokedex/internal/api"
	"github.com/Taliker/Gokedex/repl/helpers"
)

func CommandPokemon(config *Config, pokName string) error {
	if pokName == "" {
		return errors.New("please enter the name of a pokemon")
	}

	if pokemonInfo, err := api.GetPokemon(pokName, &config.Cache); err == nil {
		printPokemon(pokemonInfo)
	}
	return nil
}

func printPokemon(pokemon api.Pokemon) {
	var pf []string
	pf = append(pf, "Pokemon Info:")
	pf = append(pf, fmt.Sprintf("ID: %v", pokemon.ID))
	pf = append(pf, fmt.Sprintf("Name: %v", pokemon.Name))
	pf = append(pf, fmt.Sprintf("Base EXP: %v", pokemon.BaseExperience))
	pf = append(pf, fmt.Sprintf("Type: %v %v", pokemon.Types.Primary, pokemon.Types.Secondary))
	pf = append(pf, "Stats:")
	pf = append(pf, fmt.Sprintf("Hp:    %v", pokemon.Stats.HP))
	pf = append(pf, fmt.Sprintf("Atk:   %v", pokemon.Stats.Attack))
	pf = append(pf, fmt.Sprintf("Def:   %v", pokemon.Stats.Defense))
	pf = append(pf, fmt.Sprintf("SpAtk: %v", pokemon.Stats.SpAtk))
	pf = append(pf, fmt.Sprintf("SpDef: %v", pokemon.Stats.SpDef))
	pf = append(pf, fmt.Sprintf("Spd:   %v", pokemon.Stats.Speed))

	helpers.PrintArray(pf)
}
