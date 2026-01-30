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

	// 1️⃣ Check cache
	if data, ok := cache.Get(key); ok {
		var p PokemonResponse
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		return attemptCatch(p)
	}

	// 2️⃣ Fetch from API
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

	// 3️⃣ Parse
	var p PokemonResponse
	if err := json.Unmarshal(body, &p); err != nil {
		return err
	}

	// 4️⃣ Cache result
	cache.Add(key, body)

	// 5️⃣ Attempt catch
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
