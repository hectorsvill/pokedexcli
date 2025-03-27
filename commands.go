package main

import (
	"fmt"
	"os"
	"errors"
)

var cliCommands map[string]CliCommand

type CliCommand struct {
	name        string
	description string
	callback    func() error
	callbackWinput func(string) error
}

func (cm CliCommand) Init() {
	cliCommands = map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    Usage,
		},
		"map": {
			name:        "map",
			description: "print next locations",
			callback:    MapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "print previous locations",
			callback:    MapBack,
		},
		"explore": {
			name:        "explore",
			description: "print pokemon in location",
			callback:    Explore,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (c CliCommand) helpString() string {
	return fmt.Sprintf("%s: %s", c.name, c.description)
}

func Usage() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cliCommand := range cliCommands {
		fmt.Println(cliCommand.helpString())
	}
	return nil
}


func MapNext() error {
	if result == nil {
		locations := getLocations(locations_url)
		printLocations(locations)
	} else {
		locations := getLocations(result.Next)
		printLocations(locations)
	}
	return nil
}

func MapBack() error {
	if result.Previous == "" {
		locations := getLocations(locations_url)
		printLocations(locations)
	} else {
		locations := getLocations(result.Previous)
		printLocations(locations)
	}
	return nil
}


func Explore() error {
	if len(InputArr) != 2 {
		return errors.New("Explore(): input error")
	}
	pokemons := getLocation(InputArr[1])
	fmt.Println()
	for _, p := range pokemons {
		fmt.Printf("- %v\n", p.Name)
	}
	return nil
}