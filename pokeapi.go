package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	result           *Result
	encountersResult *PokemonEncounterResult
	statsResult      *StatsResult
	locations_url    = "https://pokeapi.co/api/v2/location-area/"
	pokemon_url      = "https://pokeapi.co/api/v2/pokemon/"
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

func getStats(pokemon string) []Stat {
	mux := &sync.RWMutex{}
	pokemon_url = pokemon_url + pokemon
	if PCache.Exist(pokemon_url) {
		entry, err := PCache.Get(pokemon_url, mux)
		
		err = json.Unmarshal([]byte(entry.val), &statsResult)
		if err != nil {
			log.Fatal(err)
		}
		return statsResult.getStats()
	} else {
		resp, err := http.Get(pokemon_url)
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

		go PCache.Add(pokemon_url, entry, mux)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal([]byte(data), &statsResult)
		if err != nil {
			log.Fatal(err)
		}

		return statsResult.getStats()
	}
}
