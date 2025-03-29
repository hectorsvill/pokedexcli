package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetLocations(url string) (Result, error) {
	res := Result{}
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &res)
		if err != nil {
			return Result{}, err
		}
		return res, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Result{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{}, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return Result{}, err
	}

	c.cache.Add(url, data)

	return res, err
}

func (c Client) GetLocation(location string) ([]Pokemon, error) {
	res := PokemonEncounterResult{}
	location = LocationsUrl + location
	if val, ok := c.cache.Get(location); ok {
		err := json.Unmarshal(val, &res)
		if err != nil {
			return nil, err
		}

		return res.getPokemon(), nil
	}

	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	c.cache.Add(location, data)

	return res.getPokemon(), nil
}
