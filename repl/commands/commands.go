package commands

// Command is a struct that represents a command that can be executed in the REPL.
type Command struct {
	Name        string
	Description string
	Callback    func() error
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
	}
}
