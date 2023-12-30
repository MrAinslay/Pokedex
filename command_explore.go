package main

import "fmt"

func commandExplore(cfg *config, s string) error {
	encountersResp, err := cfg.pokeapiClient.ListPokemon(s)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", s, "...\nFound Pokemon:")

	for _, enc := range encountersResp.PokemonEncounters {
		fmt.Println("- ", enc.Pokemon.Name)
	}
	return err
}
