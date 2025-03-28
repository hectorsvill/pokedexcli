package pokecache

import (
	"sync"
	"time"
)

var PCache PokeCache

type PokeCache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (pk PokeCache) Init() {
	PCache = PokeCache{
		cache: make(map[string]cacheEntry),
	}
}

func (pk PokeCache) Add(url string, entry cacheEntry, mux *sync.RWMutex) {
	mux.RLock()
	PCache.cache[url] = entry
	mux.RUnlock()
}

func (pk PokeCache) Get(url string, mux *sync.RWMutex) (cacheEntry, error) {
	mux.RLock()
	entry := PCache.cache[url]
	mux.RUnlock()
	return entry, nil
}

func (pk PokeCache) Exist(url string) bool {
	_, ok := PCache.cache[url]
	if ok {
		return true
	}
	return false
}
