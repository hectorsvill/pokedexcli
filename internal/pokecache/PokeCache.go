package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewPokeCache() PokeCache {
	return PokeCache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
}

func (pk PokeCache) Add(key string, value []byte) {
	pk.mux.Lock()
	defer pk.mux.Unlock()
	pk.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (pk PokeCache) Get(key string) ([]byte, bool) {
	pk.mux.Lock()
	defer pk.mux.Unlock()
	entry, ok := pk.cache[key]
	return entry.val, ok
}

func (pk PokeCache) Exist(url string) bool {
	_, ok := pk.cache[url]
	if ok {
		return true
	}
	return false
}
