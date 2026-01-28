package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
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
			fmt.Println("Your command was:", commands[0])
		} else {
			break
		}
	}
}

