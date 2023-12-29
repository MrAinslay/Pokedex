package main

import (
	"encoding/json"
	"fmt"
)

func commandMap() error {
	type Location struct {
		Count    int     `json:"count"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	locations := Location{}
	if locations.Next != nil {
		err := json.Unmarshal(getAPI(*locations.Next), &locations)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := json.Unmarshal(getAPI(""), &locations)
		if err != nil {
			fmt.Println(err)
		}

	}

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}
