package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// pokedexcli()
	str1 := "  foo   \bar     baz   "
	fields1 := strings.Fields(str1)
	fmt.Println(len(fields1))
	fmt.Printf("Fields: %q\n", fields1) // Output: Fields: ["foo" "bar" "baz"]


}

func pokedexcli() {
	initCommands()
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
