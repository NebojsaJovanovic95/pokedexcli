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
	callback func([]string) error
}

func commandExit(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Exiting program")
}

func printHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range CliMap {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
}

func commandHelp(args []string) error {
	printHelp()
	return nil
}

func commandMap(args []string) error {
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

func commandMapb(args []string) error {
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


func commandExplore(args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: explore <area_name>")
		return nil
	}

	areaName := args[0]
	return Explore(areaName)
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
		"explore": {
			name: "explore",
			description: "explores are",
			callback: commandExplore,
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
		if !more {
			break
		}
		line = scanner.Text();
		commands := cleanInput(line)
		if len(commands) == 0 {
			continue
		}

		cmdName := commands[0]
		args := commands[1:]

		command, ok := CliMap[cmdName]
		if !ok {
			fmt.Println("Unknown command:", cmdName)
			continue
		}

		if err := command.callback(args); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

