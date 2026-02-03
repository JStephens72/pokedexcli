package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon := cfg.caughtPokemon[name]
	if pokemon.Name == "" {
		return errors.New("you have not caught that pokemon")
	}

	pokemonStats := make(map[string]int)

	for _, stat := range pokemon.Stats {
		val := stat.BaseStat
		key := stat.Stat.Name
		pokemonStats[key] = val
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Print("Stats:\n")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, movementType := range pokemon.Types {
		fmt.Printf("  - %s\n", movementType.Type.Name)
	}
	/*
		fmt.Printf("  -hp: %d\n", pokemonStats["hp"])
		fmt.Printf("  -attack: %d\n", pokemonStats["attack"])
		fmt.Printf("  -defense: %d\n", pokemonStats["defense"])
		fmt.Printf("  -special-attack: %d\n", pokemonStats["special-attack"])
		fmt.Printf("  -special-defense: %d\n," pokemonStats["special-defense"])
		fmt.Printf("  -speed: %d\n", pokemonStats.("speed"))
		fmt.Print("Types:")
		fmt.Print
	*/
	return nil
}
