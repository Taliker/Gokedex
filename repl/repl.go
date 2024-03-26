package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands "github.com/Taliker/Gokedex/repl/commands"
)

func StartREPL(config *commands.Config) {
	// Start the REPL
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print(">> ")
		scanner.Scan()

		input := sanitizeInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		command, exists := commands.GetCommands()[commandName]
		if !exists {
			fmt.Println("Command not found. Type 'help' for a list of commands.")
			continue
		}

		var arg string
		if command.NeedsArg && len(input) <= 1 {
			fmt.Println("Please provide an argument for the command.")
			continue
		}

		if command.NeedsArg {
			arg = input[1]
		}

		err := command.Callback(config, arg)
		if err != nil {
			fmt.Println("An error occurred:", err)
		}
	}
}

func sanitizeInput(input string) []string {
	toLower := strings.ToLower(input)
	output := strings.Fields(toLower)
	return output
}
