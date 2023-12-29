package main

import (
	"encoding/json"
	"fmt"
)

var nextLocation string
var lastLocation string
var pntrNextLocation *string
var pntrLastLocation *string

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
			if locations.Previous != nil {
				*pntrLastLocation = *locations.Previous
			}
			return locations
		}
	}
	if pntrLastLocation != nil && *pntrLastLocation != "" {
		err := json.Unmarshal(getAPI(*pntrLastLocation), &locations)
		if err != nil {
			fmt.Println(err)
		}
		if locations.Previous != nil {
			*pntrLastLocation = *locations.Previous
			return locations
		}
		fmt.Println("There are no previous locations")
		return zeroVal
	}
	fmt.Println("You have not visited any other locations yet")
	return zeroVal
}
