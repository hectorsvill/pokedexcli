package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func removeDoubleSpace(str string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(str, " ")
}

func cleanInput(text string) []string {
	str := strings.ToLower(text)
	str = strings.Trim(str, " ")
	str = removeDoubleSpace(str)
	return strings.Split(str, " ")
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func (c cliCommand) helpString() string {
	return fmt.Sprintf("%s: %s", c.name, c.description)
}

func usage(cliCommands map[string]cliCommand) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cliCommand := range cliCommands {
		fmt.Println(cliCommand.helpString())
	}
	return nil
}




type Result struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
}


var (
	result *Result
)



func getLocations(url string) []Location {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error getting location: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}
	
	return result.Results
}
func printLocations(locations []Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func Map(action string) {
	url := "https://pokeapi.co/api/v2/location-area/"

	if action == "next" {
		if  result ==  nil {
			locations := getLocations(url)
			printLocations(locations)
		} else {
			locations := getLocations(result.Next)
			printLocations(locations)
		}
	} else {
		if  result.Previous == "" {
			locations := getLocations(url)
			printLocations(locations)
		} else {
			locations := getLocations(result.Previous)
			printLocations(locations)
		}
	}
}




func pokedexcli() {
	cliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "print all locations",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedexcli > ")
		scanner.Scan()
		str := scanner.Text()
		inputArr := cleanInput(str)
		command := inputArr[0]
		switch command {
		case "exit":
			commandExit()
		case "map":
			Map("next")
		case "mapb":
			Map("")
		default:
			usage(cliCommands)
		}
	}
}

func main() {
	pokedexcli()

}
