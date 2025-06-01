package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Dirza1/Pokedex/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    nil,
		},
		"map": {
			name:        "map",
			description: "displays 20 pokemon locations",
			callback:    commandMap,
		},
	}
	help := commands["help"]
	help.callback = commandHelp(commands)
	commands["help"] = help

	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		input := scanner.Text()
		returned := cleanInput(input)
		_, ok := commands[returned[0]]
		if !ok {
			fmt.Print("Unknown command\n")
		} else {
			commands[returned[0]].callback()
		}

	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Fields(lower)
	return split
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) func() error {
	return func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")

		for _, value := range commands {
			fmt.Printf("%s: %s\n", value.name, value.description)
		}
		return nil
	}
}

func commandMap() error {
	/*locations :=*/ pokeapi.Main()
	/*for location := range locations {
		fmt.Println(location)
	}*/
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
