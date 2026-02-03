package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	respPokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(651)

	fmt.Printf("Throwing a Pokeball at %s...\n", respPokemon.Name)
	fmt.Printf("Base Experience for %s: %d\n", respPokemon.Name, respPokemon.BaseExperience)
	fmt.Printf("You rolled a %d\n", res)
	if respPokemon.BaseExperience > res {
		fmt.Printf("%s escaped!\n", respPokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", respPokemon.Name)

	cfg.caughtPokemon[respPokemon.Name] = respPokemon
	return nil
}
