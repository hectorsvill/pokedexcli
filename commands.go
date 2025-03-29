package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
	// "github.com/hectorsvill/pokedexcli/internal/pokeapi"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCliCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
		"inspect": {
			name:        "inspect",
			description: "print pokemon stats",
			callback:    Inspect,
		},
		"catch": {
			name:        "catch",
			description: "try catching a pokemon",
			callback:    Catch,
		},
	}
}

func (cmd CliCommand) helpString() string {
	return fmt.Sprintf("%s: %s", cmd.name, cmd.description)
}

func commandExit(cfg *config) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func Usage(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cliCommand := range getCliCommands() {
		fmt.Println(cliCommand.helpString())
	}
	return nil
}

func MapNext(cfg *config) error {
	if len(cfg.inputArr) != 1 {
		panic("input error")
	}

	result := cfg.client.GetLocations(cfg.nextLocation)

	cfg.nextLocation = result.Next
	cfg.previousLocation = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func MapBack(cfg *config) error {
	if len(cfg.inputArr) != 1 {
		panic("input error")
	}

	result := cfg.client.GetLocations(cfg.previousLocation)

	cfg.nextLocation = result.Next
	cfg.previousLocation = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func Explore(cfg *config) error {
	if len(cfg.inputArr) != 2 {
		return errors.New("Explore(): input error")
	}
	pokemons := cfg.client.GetLocation(cfg.inputArr[1])
	fmt.Println()
	for _, p := range pokemons {
		fmt.Printf("- %v\n", p.Name)
	}
	return nil
}

func Inspect(cfg *config) error {
	if len(cfg.inputArr) != 2 {
		return errors.New("Inspect(): input error")
	}

	stats,_ := cfg.client.GetStats(cfg.inputArr[1])
	for _, stat := range stats {
		fmt.Printf("  -%v: %v\n", stat.Name, stat.Base_Stat)
	}

	return nil
}

func Catch(cfg *config) error {
	if len(cfg.inputArr) != 2 {
		return errors.New("Catch(): input error\n")
	}
	
	pokemon := cfg.inputArr[1]
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

	stats, err := cfg.client.GetStats(pokemon)
	if err != nil {
		fmt.Printf("\n%v is not a pokemon...\n", pokemon)
		return nil
	}

	hpBaseStat := stats[0].Base_Stat
	defenseStat := stats[2].Base_Stat
	
	time.Sleep(500 * time.Millisecond)
	
	randVal := rand.Intn(hpBaseStat)
	catchRate := int(float64( hpBaseStat/2) + (0.10 * float64(defenseStat)))

	if randVal > catchRate  {
		fmt.Printf("%v was caught!\n", pokemon)
	} else {
		fmt.Printf("%v escaped!\n", pokemon)
	}

	return nil
}
