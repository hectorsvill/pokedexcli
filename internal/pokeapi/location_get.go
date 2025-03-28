package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetLocations(url string) Result {
	res := Result{}
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &res)
		if err != nil {
			panic(err)
		}
		return res
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}

	c.cache.Add(url, data)

	return res
}
