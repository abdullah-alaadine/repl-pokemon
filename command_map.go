package main

import (
	"fmt"
	"log"

	"github.com/abdullah-alaadine/repl-pokemon/internal/pokeapi"
)

func callbackMap() error {
	httpClient := pokeapi.NewClient()
	locationAreas, err := httpClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, area := range locationAreas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}
