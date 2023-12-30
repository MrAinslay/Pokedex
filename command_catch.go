package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, s string) error {
	n := rand.Intn(341)
	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(s)
	if err != nil {
		return err
	}
	pokemonExperience := pokemon.BaseExperience
	if pokemonExperience > n {
		fmt.Println(pokemon.Name, "escaped!")
		return nil
	}
	fmt.Println(pokemon.Name, "was caught!")

	return nil
}
