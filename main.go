package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

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

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	cliMap := map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
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
			if command, ok := cliMap[commands[0]]; ok {
				command.callback()
			} else {
				fmt.Println("Error: ", err)
			}
		} else {
			break
		}
	}
}

