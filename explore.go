package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

)

func Explore(areaName string) error {
	key := "explore:" + strings.ToLower(areaName)

	// check cache first
	if data, ok := cache.Get(key); ok {
		var names []string
		if err := json.Unmarshal(data, &names); err != nil {
			return err
		}
		printPokemon(areaName, names)
		return nil
	}

	// fetch from API
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", areaName)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var res LocationAreaDetail
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	// extract names
	names := []string{}
		for _, encounter := range res.PokemonEncounters {
		names = append(names, encounter.Pokemon.Name)
	}

	// store in cache
	cache.Add(key, mustMarshal(names))

	printPokemon(areaName, names)
	return nil
}

// helper to print nicely
func printPokemon(area string, names []string) {
	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, name := range names {
		fmt.Println(" -", name)
	}
}

// helper to marshal without error
func mustMarshal(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}

