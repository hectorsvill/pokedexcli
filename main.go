package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func init() {
	CliCommand{}.Init()
	PokeCache{}.Init()
}

func main() {
	pokedexcli()
	block := make(chan struct{})
	<-block
}

func pokedexcli() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedexcli > ")
		scanner.Scan()
		str := scanner.Text()
		inputArr := strings.Fields(str)
		command := inputArr[0]
		if cmd, ok := cliCommands[command]; ok {
			if err := cmd.callback(); err != nil {
				log.Println("Error", err)
			}
		} else {
			Usage()
		}
	}
}
