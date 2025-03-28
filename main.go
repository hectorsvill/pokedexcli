package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hectorsvill/pokedexcli/internal/pokeapi"
)

type config struct {
	inputArr         []string
	client           pokeapi.Client
	nextLocation     string
	previousLocation string
}

func main() {
	cfg := &config{
		inputArr:     []string{},
		client:       pokeapi.NewClient(5 * time.Second),
		nextLocation: pokeapi.LocationsUrl,
	}

	pokedexcli(cfg)
}

func pokedexcli(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	CliCommands := getCliCommands()

	for {
		fmt.Print("pokedexcli > ")
		scanner.Scan()

		str := scanner.Text()
		str = strings.ToLower(str)
		cfg.inputArr = strings.Fields(str)

		command := cfg.inputArr[0]

		if cmd, ok := CliCommands[command]; ok {
			if err := cmd.callback(cfg); err != nil {
				log.Println("Error", err)
			}
		} else {
			Usage(cfg)
		}
	}
}
