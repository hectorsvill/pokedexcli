package pokeapi

type PokemonEncounterResult struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (per PokemonEncounterResult) getPokemon() []Pokemon {
	pokemons := []Pokemon{}
	for _, p := range per.PokemonEncounters {
		pokemons = append(pokemons, p.Pokemon)
	}
	return pokemons
}
