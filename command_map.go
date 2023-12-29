package main

import (
	"fmt"
)

func commandMap() error {
	locations := getLocations(0)
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}
