package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getAPI(s string) []byte {
	link := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", s)
	res, err := http.Get(link)
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