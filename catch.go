package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"math/rand"
)

func CatchPokemon(name string) error {
	key := "pokemon:" + strings.ToLower(name)

	if data, ok := cache.Get(key); ok {
		var p PokemonResponse
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		return attemptCatch(p)
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("pokemon %s not found", name)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var p PokemonResponse
	if err := json.Unmarshal(body, &p); err != nil {
		return err
	}

	cache.Add(key, body)

	return attemptCatch(p)
}

func attemptCatch(p PokemonResponse) error {
	chance := rand.Intn(100)

	if chance > p.BaseExperience {
		fmt.Printf("You caught %s!\n", p.Name)
	} else {
		fmt.Printf("%s escaped!\n", p.Name)
	}

	return nil
}
