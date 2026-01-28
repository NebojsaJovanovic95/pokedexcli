package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"net/http"
	"encoding/json"
)

var CliMap map[string]cliCommand

type cliCommand struct {
	name string
	description string
	callback func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Exiting program")
}

func printHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, value := range CliMap {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
}

func commandHelp() error {
	printHelp()
	return nil
}

type LocationAreaResponce struct {
	Results []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func commandMap() error {
	url := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch location areas")
	}
	var data LocationAreaResponce
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	} else {
		for _, location := range data.Results {
			fmt.Println(location.Name)
		}
	}
	return nil
}

func commandMapb() error {
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	CliMap = map[string]cliCommand{ 
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "loads list of locations",
			callback: commandMap,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	var line string;
	for ;; {
		fmt.Print("Pokedex > ")
		more := scanner.Scan();
		err := scanner.Err();
		if err != nil {
			fmt.Println("Error: %w", err)
		}
		if more {
			line = scanner.Text();
			commands := cleanInput(line)
			if command, ok := CliMap[commands[0]]; ok {
				command.callback()
			} else {
				fmt.Println("Error: ", err)
			}
		} else {
			break
		}
	}
}

