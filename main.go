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
	config := &config{}
	config.Next = nil
	config.Previous = nil
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    nil,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    nil,
		},
		"map": {
			name:        "map",
			description: "displays 20 pokemon locations",
			callback:    nil,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the 20 previous locations",
			callback:    nil,
		},
	}
	help := commands["help"]
	help.callback = commandHelp(commands, config)
	commands["help"] = help

	exit := commands["exit"]
	exit.callback = commandExit(config)
	commands["exit"] = exit

	mapss := commands["map"]
	mapss.callback = commandMap(config)
	commands["map"] = mapss

	mapb := commands["mapb"]
	mapb.callback = commandMapB(config)
	commands["mapb"] = mapb

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

func commandExit(config *config) func() error {
	return func() error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
		return nil
	}
}

func commandHelp(commands map[string]cliCommand, config *config) func() error {
	return func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")

		for _, value := range commands {
			fmt.Printf("%s: %s\n", value.name, value.description)
		}
		return nil
	}
}

func commandMap(config *config) func() error {
	return func() error {
		var maps pokeapi.Maps
		if config.Next != nil {
			maps = pokeapi.Main(*config.Next)
		} else {
			maps = pokeapi.Main("https://pokeapi.co/api/v2/location-area")
			fmt.Printf("Next type: %T\n", maps.Next)
		}
		for _, location := range maps.Results {
			fmt.Println(location.Name)
		}
		config.Next = maps.Next
		config.Previous = maps.Previous
		return nil
	}
}

func commandMapB(config *config) func() error {
	return func() error {
		if config.Previous != nil {
			maps := pokeapi.Main(*config.Previous)

			for _, location := range maps.Results {
				fmt.Println(location.Name)
			}
			config.Next = maps.Next
			config.Previous = maps.Previous
		} else {
			fmt.Println("you're on the first page")
		}

		return nil
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next     *string
	Previous *string
}
