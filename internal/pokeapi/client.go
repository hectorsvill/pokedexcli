package pokeapi

import (
	"net/http"
	"time"

	"github.com/hectorsvill/pokedexcli/internal/pokecache"

)

type Client struct {
	cache pokecache.PokeCache
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewPokeCache(),
		httpClient: http.Client {
			Timeout: timeout,
		},
	}
}
