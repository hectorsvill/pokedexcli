package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client)GetLocations(url string) []Location {
	locations := []Location{}
	if val, ok := c.pokecache.Get(url); ok {
		err := json.Unmarshal(val, &locations)
		if err != nil {
			panic(err)
		}
		return locations
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		panic(err)
	}

	go c.pokecache.Add(url, data)
	
	return locations
}