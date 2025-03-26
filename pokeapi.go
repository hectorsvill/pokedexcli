package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
}

var (
	result *Result
)

func getLocations(url string) []Location {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error getting location: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result.Results
}
func printLocations(locations []Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func MapNext() error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if result == nil {
		locations := getLocations(url)
		printLocations(locations)
	} else {
		locations := getLocations(result.Next)
		printLocations(locations)
	}

	return nil
}

func MapBack() error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if result.Previous == "" {
		locations := getLocations(url)
		printLocations(locations)
	} else {
		locations := getLocations(result.Previous)
		printLocations(locations)
	}
	return nil
}
