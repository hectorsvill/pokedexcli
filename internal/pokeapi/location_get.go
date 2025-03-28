package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client)GetLocations(url string) []Location {
	
	if val, ok := c.pokecache.Get(url); ok {
		locations := []Location{}
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

	result := Result{}

	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	c.pokecache.Add(url, data)
	
	return result.Results
}