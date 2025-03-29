pokedexcli is a Pokedex in a command-line REPL using the [PokeAPI](https://pokeapi.co/).

### Project features
- HTTP requests
- parse JSON 
- [Cache struct](https://github.com/hectorsvill/pokedexcli/blob/main/internal/pokecache/PokeCache.go) to hold a map[string]cacheEntry and a mutex to protect the map across goroutines

#### Download and run:
```bash
git clone https://github.com/hectorsvill/pokedexcli.git
cd pokedexcli
go run .
```
#### Example usage: 
##### help
Get usage information
```bash
pokedexcli > help
Welcome to the Pokedex!
Usage:

mapb: print previous locations
explore <location>: print pokemon in location
inspect <pokemon>: print pokemon stats
catch <pokemon>: try catching a pokemon
pokedex: get my pokemon
exit: Exit the Pokedex
help: Displays a help message
map: print next locations
```
##### map
Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map will display the next 20 locations.
```bash
pokedexcli > map
canalave-city-area
.
.
```
##### mabb
Displays the names of 20 previous location areas in the Pokemon world. Each subsequent call to mapb will display the next 20 locations.
```bash
pokedexcli > mapb
canalave-city-area
.
.
```

##### explore <location>
Takes the name of a location and print pokemon located there
```bash
pokedexcli > explore pastoria-city-area

- tentacool
- tentacruel
- magikarp
- gyarados
- remoraid
- octillery
- wingull
- pelipper
- shellos
- gastrodon
```

##### catch <pokemon>
takes the name of a Pokemon as an argument. If pokemon is caught, its stored in pokedex.
- Not using the pokemon's "base experience" to determine the chance of catching it.
```bash
pokedexcli > catch wingull  
Throwing a Pokeball at wingull...
wingull escaped!
pokedexcli > catch wingull
Throwing a Pokeball at wingull...
You may now inspect it with the inspect command.
pokedexcli > 

```

##### inspect <pokemon>
Display stats for wingull.

```bash
pokedexcli > inspect wingull
  -hp: 40
  -attack: 30
  -defense: 30
  -special-attack: 55
  -special-defense: 30
  -speed: 85

```
##### pokedex
Print pokemon in pokedex
```bash
pokedexcli > pokedex 
 - wingull
```




