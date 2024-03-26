package commands

import "github.com/Taliker/Gokedex/internal/cache"

// Command is a struct that represents a command that can be executed in the REPL.
type Command struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
	NeedsArg    bool
}

type Config struct {
	nextURL string
	prevURL string
	Cache   cache.Cache
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Prints the help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the REPL",
			Callback:    CommandExit,
		},
		"mapf": {
			Name:        "mapf",
			Description: "Prints the next locations in the Pokemon world",
			Callback:    CommandMapF,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Prints the previous locations in the Pokemon world",
			Callback:    CommandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location in the Pokemon world",
			Callback:    CommandExplore,
			NeedsArg:    true,
		},
		"pokemon": {
			Name:        "pokemon",
			Description: "Get information about a Pokemon",
			Callback:    CommandPokemon,
			NeedsArg:    true,
		},
	}
}
