package commands

// Command is a struct that represents a command that can be executed in the REPL.
type Command struct {
	name        string
	description string
	callback    func() error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Prints the help message",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the REPL",
			callback:    CommandExit,
		},
	}
}
