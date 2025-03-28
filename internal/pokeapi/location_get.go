package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (c Client)getLocations(locationsUrl string) []Location {
	mux := &sync.RWMutex{}

	if PCache.Exist(locationsUrl) {
		entry, err := PCache.Get(locationsUrl, mux)
		err = json.Unmarshal(entry.val, &result)
		if err != nil {
			panic(err)
		}
	}

	req, err := http.NewRequest("GET", locationsUrl, nil)
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
	
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       data,
	}

	go PCache.Add(locationsUrl, entry, mux)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	return result.Results
}