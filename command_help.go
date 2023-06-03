package main

import "fmt"

func callbackHelp() {
	fmt.Println("Available commands:")

	availableCommands := getCommands()
	for _, val := range availableCommands {
		fmt.Printf(" - %s: %s\n", val.name, val.description)
	}
	fmt.Println()
}
