package main

import (
	"encoding/json"
	"fmt"
)

var n int
var nextLocation string
var lastLocation string
var pntrNextLocation *string
var pntrLastLocation *string
var pntrN *int

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocations(i int) Location {
	pntrN = &n
	pntrNextLocation = &nextLocation
	pntrLastLocation = &lastLocation
	var zeroVal Location

	locations := Location{}
	if i == 0 {
		if pntrNextLocation != nil {
			err := json.Unmarshal(getAPI(*pntrNextLocation), &locations)
			if err != nil {
				fmt.Println(err)
			}
			*pntrNextLocation = *locations.Next
			*pntrN++
			if locations.Previous != nil {
				*pntrLastLocation = *locations.Previous
			}
			return locations
		}
	}
	if pntrLastLocation != nil && *pntrLastLocation != "" {
		if n == 1 {
			fmt.Println("There are no previous locations")
			return zeroVal
		}
		err := json.Unmarshal(getAPI(*pntrLastLocation), &locations)
		if err != nil {
			fmt.Println(err)
		}
		*pntrN--
		if locations.Previous != nil {
			*pntrLastLocation = *locations.Previous
			return locations
		}
		*pntrNextLocation = *locations.Next
		return locations
	}
	fmt.Println("You have not visited any other locations yet")
	return zeroVal
}
