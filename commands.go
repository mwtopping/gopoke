package main

import (
	"encoding/json"
	"fmt"
	"gopoke/internal/pokeapi"
	"math/rand"
	"os"
)

func commandExit(config *Config, s string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, s string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("CLI Usage:")

	Coms := Commands()

	for _, Com := range Coms {
		fmt.Printf("%v: %v\n", Com.name, Com.description)
	}
	return nil
}

func commandMapf(config *Config, s string) error {

	// check if data exists in cache
	cache := config.pokeCache

	url := "https://pokeapi.co/api/v2/location-area/"

	if config.Next == nil {
		url = url
	} else if val, ok := cache.Get(*config.Next); ok {
		fmt.Println("Exists in cache")
		locs := pokeapi.PokeResponse{}
		err := json.Unmarshal(val.Val, &locs)
		if err != nil {
			fmt.Println("Error unmarshalling")
			fmt.Println(err)
			return err
		}
		for _, loc := range locs.Results {
			fmt.Println(loc.Name)
		}
		config.Next = locs.Next
		config.Previous = locs.Previous
		return nil

	} else {
		fmt.Println("Doesn't exist in cache")
		url = *config.Next
	}

	locs, err := config.pokeClient.GetLocations(&url)
	if err != nil {
		return err
	}
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	rawdata, rawerr := json.Marshal(locs)
	if rawerr != nil {
		return rawerr
	}
	cache.Add(url, rawdata)

	config.Next = locs.Next
	config.Previous = locs.Previous

	return nil
}

func commandMapb(config *Config, s string) error {

	if config.Previous == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	cache := config.pokeCache

	url := "https://pokeapi.co/api/v2/location-area/"

	if val, ok := cache.Get(*config.Previous); ok {
		fmt.Println("Exists in cache")
		locs := pokeapi.PokeResponse{}
		err := json.Unmarshal(val.Val, &locs)
		if err != nil {
			fmt.Println("Error unmarshalling")
			fmt.Println(err)
			return err
		}
		for _, loc := range locs.Results {
			fmt.Println(loc.Name)
		}
		config.Next = locs.Next
		config.Previous = locs.Previous
		return nil

	} else {
		fmt.Println("Doesn't exist in cache")
		url = *config.Previous
	}

	locs, err := config.pokeClient.GetLocations(&url)
	if err != nil {
		return err
	}
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	rawdata, rawerr := json.Marshal(locs)
	if rawerr != nil {
		return rawerr
	}
	cache.Add(url, rawdata)

	config.Next = locs.Next
	config.Previous = locs.Previous

	return nil
}

func commandExplore(config *Config, loc string) error {

	// check if data exists in cache
	cache := config.pokeCache

	url := "https://pokeapi.co/api/v2/location-area/" + loc

	if val, ok := cache.Get(url); ok {
		fmt.Println("Exists in cache")
		locs := pokeapi.PokeLocation{}
		err := json.Unmarshal(val.Val, &locs)
		if err != nil {
			fmt.Println("Error unmarshalling")
			fmt.Println(err)
			return err
		}
		for _, loc := range locs.PokemonEncounters {
			fmt.Println(loc.Pokemon.Name)
		}
		return nil

	} else {
		fmt.Println("Doesn't exist in cache")
	}

	locs, err := config.pokeClient.GetPokemon(&url)
	if err != nil {
		return err
	}

	for _, loc := range locs.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}

	rawdata, rawerr := json.Marshal(locs)
	if rawerr != nil {
		return rawerr
	}
	fmt.Println("Addind", url, "to cache")
	cache.Add(url, rawdata)

	return nil
}

func commandCatch(config *Config, mon string) error {

	// check if data exists in cache
	cache := config.pokeCache

	url := "https://pokeapi.co/api/v2/pokemon/" + mon

	if val, ok := cache.Get(url); ok {
		fmt.Println("Exists in cache")
		locs := pokeapi.Pokemon{}
		err := json.Unmarshal(val.Val, &locs)
		if err != nil {
			fmt.Println("Error unmarshalling")
			fmt.Println(err)
			return err
		}

		fmt.Printf("Throwing a Pokeball at %v...\n", mon)
		catchrate := float32(255-locs.BaseExperience) / 256.0
		if rand.Float32() > catchrate {
			fmt.Println(mon, "was caught!")
			config.pokemon[mon] = locs
		} else {
			fmt.Println(mon, "got away")
		}

		return nil

	} else {
		fmt.Println("Doesn't exist in cache")
	}

	locs, err := config.pokeClient.GetPokemonStats(&url)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", mon)
	catchrate := float32(255-locs.BaseExperience) / 256.0
	if rand.Float32() > catchrate {
		fmt.Println(mon, "was caught!")
		config.pokemon[mon] = locs
	} else {
		fmt.Println(mon, "got away")
	}

	rawdata, rawerr := json.Marshal(locs)
	if rawerr != nil {
		return rawerr
	}
	fmt.Println("Addind", url, "to cache")
	cache.Add(url, rawdata)

	return nil
}

func commandInspect(config *Config, s string) error {

	if val, ok := config.pokemon[s]; ok {
		fmt.Printf("Name: %v\n", val.Name)
		fmt.Printf("Weight: %v\n", val.Weight)
		fmt.Printf("Height: %v\n", val.Height)
		fmt.Printf("Types:\n")
		for _, v := range config.pokemon[s].Types {
			fmt.Printf(" -%v\n", v.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}

func commandPokedex(config *Config, s string) error {

	fmt.Println("Your Pokedex:")
	for _, poke := range config.pokemon {
		fmt.Printf(" - %v\n", poke.Name)
	}

	return nil
}
