package main

import (
	"encoding/json"
	"fmt"
	"gopoke/internal/pokeapi"
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
