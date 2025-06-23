package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	textL := strings.ToLower(text)
	newInput := strings.Fields(textL)

	return newInput
}
func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	var cfg config
	cfg = config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
	}
	cfgPtr := &cfg
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfgPtr)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Uknown command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpMessage,
		},
		"map": {
			name:        "map",
			description: "Displays 20 cities around canalave-city",
			callback:    getMap,
		},
		"mapb": {
			name:        "previousmap",
			description: "Displays previous 20 cities around canalave-city",
			callback:    previousMap,
		},
	}
}
