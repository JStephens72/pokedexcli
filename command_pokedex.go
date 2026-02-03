package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {

	myPokedex := cfg.caughtPokemon
	if len(myPokedex) == 0 {
		return errors.New("your pokedex is empty")
	}

	fmt.Print("Your Pokedex:\n")
	for _, pokemon := range myPokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
