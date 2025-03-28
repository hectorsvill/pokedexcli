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

func (pk PokeCache) New() PokeCache {
	return PokeCache {
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
}

func (pk PokeCache) Add(key string, value []byte) {
	pk.mux.Lock()
	defer pk.mux.Unlock()
	PCache.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: value, 
	}
}

func (pk PokeCache) Get(key string) ([]byte, bool) {
	pk.mux.Lock()
	defer pk.mux.Unlock()
	entry, ok := PCache.cache[key]
	return entry.val, ok
}

func (pk PokeCache) Exist(url string) bool {
	_, ok := PCache.cache[url]
	if ok {
		return true
	}
	return false
}
