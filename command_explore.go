package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no location provided")
	}
	respPokemons, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	for _, p := range respPokemons.PokemonEncounters {
		fmt.Println(p.Pokemon.Name)
	}

	return nil
}
