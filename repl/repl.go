package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	commands "github.com/Taliker/Gokedex/repl/commands"
)

func StartREPL() {
	// Start the REPL
	scanner := bufio.NewScanner(os.Stdin)
	var config = commands.Config{}
	config.Cache = config.Cache.NewCache(5 * time.Minute)
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

		err := command.Callback(&config)
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
