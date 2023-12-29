package main

import (
	"time"

	"github.com/MrAinslay/Pokedex/internal/pokeapi"

	"github.com/MrAinslay/Pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	pokeCache := pokecache.NewCache(5 * time.Second)
	go pokeCache.ReapLoop()
	startRpl(cfg)
}
