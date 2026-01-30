package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"time"
	"pokedexcli/internal/pokecache"
)

var cache *pokecache.Cache // global cache

func init() {
	// Initialize the cache with 5-second expiry
	cache = pokecache.NewCache(5 * time.Second)
}

var CliMap map[string]cliCommand
var LAP LocationAreaPaginator;

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

func commandMap() error {
	laList, err := LAP.NextPage()
	if err != nil {
		return err
	} else {
		for _, la := range laList {
			fmt.Println(la.Name)
		}
		return nil
	}
}

func commandMapb() error {
	laList, err := LAP.PrevPage()
	if err != nil {
		return err
	} else {
		for _, la := range laList {
			fmt.Println(la.Name)
		}
		return nil
	}
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
		"mapb": {
			name: "mapb",
			description: "loads previous list of locations",
			callback: commandMapb,
		},
}
	LAP.Init(20)
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

