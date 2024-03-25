package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands "github.com/Taliker/Gokedex/repl/commands"
)

func StartREPL() {
	// Start the REPL
	scanner := bufio.NewScanner(os.Stdin)
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
			fmt.Println("Command not found.")
			continue
		}

		err := command.Callback()
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
