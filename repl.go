package main

import (
	"bufio"
	"fmt"
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
	}
	return myCommands
}
