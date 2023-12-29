package main

import (
	"io"
	"log"
	"net/http"
)

func getAPI(s string) []byte {
	if s == "" {
		s = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(s)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return body
}
