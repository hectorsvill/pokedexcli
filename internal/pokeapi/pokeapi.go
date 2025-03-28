package pokeapi

const (
	LocationsUrl = "https://pokeapi.co/api/v2/location-area/"
	pokemonUrl   = "https://pokeapi.co/api/v2/pokemon/"
)

// func getLocations(locationsUrl string) []Location {
// 	mux := &sync.RWMutex{}
// 	if PCache.Exist(locationsUrl) {
// 		entry, err := PCache.Get(locationsUrl, mux)
// 		err = json.Unmarshal(entry.val, &result)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	resp, err := http.Get(locationsUrl)
// 	if err != nil {
// 		log.Fatalf("error getting location: %v", err)
// 	}
// 	defer resp.Body.Close()
// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	entry := CacheEntry{
// 		createdAt: time.Now(),
// 		val:       data,
// 	}
// 	go PCache.Add(locationsUrl, entry, mux)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = json.Unmarshal(data, &result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return result.Results
// }

// func printLocations(locations []Location) {
// 	for _, location := range locations {
// 		fmt.Println(location.Name)
// 	}
// }

// func getLocation(location string) []Pokemon {
// 	location = locationsUrl + location
// 	mux := &sync.RWMutex{}
// 	if PCache.Exist(location) {
// 		entry, err := PCache.Get(location, mux)

// 		err = json.Unmarshal(entry.val, &encountersResult)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return encountersResult.getPokemon()
// 	}

// 	resp, err := http.Get(location)
// 	if err != nil {
// 		log.Fatalf("error getting location: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	entry := CacheEntry{
// 		createdAt: time.Now(),
// 		val:       data,
// 	}

// 	go PCache.Add(location, entry, mux)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = json.Unmarshal([]byte(data), &encountersResult)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//fmt.Println(encountersResult.PokemonEncounters[0].Pokemon.Name)
// 	return encountersResult.getPokemon()
// }

// func getStats(pokemon string) []Stat {
// 	mux := &sync.RWMutex{}
// 	pokemonUrl = pokemonUrl + pokemon
// 	if PCache.Exist(pokemon) {
// 		entry, err := PCache.Get(pokemon, mux)
// 		err = json.Unmarshal([]byte(entry.val), &statsResult)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return statsResult.getStats()
// 	}

// 	resp, err := http.Get(pokemonUrl)
// 	if err != nil {
// 		log.Fatalf("error getting location: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	entry := CacheEntry{
// 		createdAt: time.Now(),
// 		val:       data,
// 	}

// 	go PCache.Add(pokemon, entry, mux)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = json.Unmarshal([]byte(data), &statsResult)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return statsResult.getStats()

// }
