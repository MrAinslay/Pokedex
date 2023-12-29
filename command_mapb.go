package main

import "fmt"

func commandMapb() error {
	locations := getLocations(1)
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}
