package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
	"errors"
)


var (
	result *Result
	encountersResult *PokemonEncounterResult
	locations_url = "https://pokeapi.co/api/v2/location-area/"
)

func getLocations(locations_url string) []Location {
	mux := &sync.RWMutex{}

	if PCache.Exist(locations_url) {
		entry, err := PCache.Get(locations_url, mux)

		err = json.Unmarshal(entry.val, &result)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		resp, err := http.Get(locations_url)
		if err != nil {
			log.Fatalf("error getting location: %v", err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		entry := CacheEntry{
			createdAt: time.Now(),
			val:       data,
		}

		go PCache.Add(locations_url, entry, mux)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &result)
		if err != nil {
			log.Fatal(err)
		}
	}
	
	return result.Results
}

func printLocations(locations []Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func MapNext() error {
	if result == nil {
		locations := getLocations(locations_url)
		printLocations(locations)
	} else {
		locations := getLocations(result.Next)
		printLocations(locations)
	}
	return nil
}

func MapBack() error {
	if result.Previous == "" {
		locations := getLocations(locations_url)
		printLocations(locations)
	} else {
		locations := getLocations(result.Previous)
		printLocations(locations)
	}
	return nil
}



func getLocation(location string) []Pokemon {
	location = locations_url + location
	mux := &sync.RWMutex{}
	if PCache.Exist(location) {
		entry, err := PCache.Get(location, mux)

		err = json.Unmarshal(entry.val, &encountersResult)
		if err != nil {
			log.Fatal(err)
		}
		return encountersResult.getPokemon()
	} else {
		resp, err := http.Get(location)
		if err != nil {
			log.Fatalf("error getting location: %v", err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		entry := CacheEntry{
			createdAt: time.Now(),
			val:       data,
		}

		go PCache.Add(location, entry, mux)
		if err != nil {
			log.Fatal(err)
		}
		
		err = json.Unmarshal([]byte(data), &encountersResult)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(encountersResult.PokemonEncounters[0].Pokemon.Name)
		return encountersResult.getPokemon()
	}
}


func Explore() error {
	if len(InputArr) != 2 {
		return errors.New("Explore(): input error")
	}
	pokemons := getLocation(InputArr[1])
	for _, p := range pokemons {
		fmt.Printf("- %v\n", p.Name)
	}
	return nil
}





