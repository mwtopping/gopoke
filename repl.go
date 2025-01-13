package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("CLI Usage:")

	Coms := Commands()

	for _, Com := range Coms {
		fmt.Printf("%v: %v\n", Com.name, Com.description)
	}
	return nil
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	locs := pokeResponse{}

	err = json.Unmarshal(body, &locs)
	if err != nil {
		return err
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	return nil
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
	callback    func() error
}

func startRepl() {

	myScanner := bufio.NewScanner(os.Stdin)

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
			val.callback()
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
			description: "Provides Locations",
			callback:    commandMap},
	}
	return myCommands
}
