package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	threshold := rand.Intn(100)
	randNum := rand.Intn(pokemon.BaseExperience + 100)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
