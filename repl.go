package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopoke/internal/pokeapi"
	"gopoke/internal/pokecache"
)

type Config struct {
	pokeClient pokeapi.Client
	pokeCache  *pokecache.Cache
	Next       *string
	Previous   *string
	pokemon    map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	var mylist []string

	for _, s := range strings.Fields(text) {
		mylist = append(mylist, strings.ToLower(s))
	}

	return mylist
}

type Command struct {
	name        string
	description string
	callback    func(c *Config, s string) error
}

func startRepl(cache *pokecache.Cache) {

	myScanner := bufio.NewScanner(os.Stdin)

	pokeclient := pokeapi.NewClient()

	config := Config{pokeClient: pokeclient,
		pokeCache: cache, pokemon: make(map[string]pokeapi.Pokemon)}

	// main loop
	for {
		fmt.Print("Pokedex >")
		myScanner.Scan()
		t := myScanner.Text()

		words := cleanInput(t)
		if len(words) == 0 {
			continue
		}

		if val, ok := Commands()[words[0]]; ok {
			if len(words) > 1 {
				val.callback(&config, words[1])
			} else {
				val.callback(&config, "")
			}
		} else {
			fmt.Println("Unknown Command")
		}

	}

}

func Commands() map[string]Command {
	// command definitions
	myCommands := map[string]Command{
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit},
		"help": {
			name:        "help",
			description: "Lists possible commands",
			callback:    commandHelp},
		"map": {
			name:        "map",
			description: "Provides next 20 locations",
			callback:    commandMapf},
		"bmap": {
			name:        "mapb",
			description: "Provides previous 20 locations",
			callback:    commandMapb},
		"explore": {
			name:        "explore",
			description: "List available pokemon in a locatin",
			callback:    commandExplore},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch},
		"inspect": {
			name:        "inspect",
			description: "Read Pokedex entry",
			callback:    commandInspect},
		"pokedex": {
			name:        "pokedex",
			description: "List pokedex",
			callback:    commandPokedex},
	}
	return myCommands
}
