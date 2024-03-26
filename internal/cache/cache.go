package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mux     *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entries: make(map[string]CacheEntry),
		mux:     &sync.Mutex{},
	}
	go cache.readLoop(interval)
	return newCache
}

func (cache *Cache) Add(key string, val []byte) {
	var entry CacheEntry = CacheEntry{createdAt: time.Now(),
		val: val,
	}
	cache.mux.Lock()

	cache.entries[key] = entry
	defer cache.mux.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()

	entry, ok := cache.entries[key]
	defer cache.mux.Unlock()
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache) readLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		if len(cache.entries) == 0 {
			continue
		}
		fmt.Printf("Cleaning cache, Lenght:%d", len(cache.entries))
		cache.mux.Lock()
		cache.mux.Lock()
		for key, entry := range cache.entries {
			if time.Since(entry.createdAt) > interval {
				delete(cache.entries, key)
			}
		}
		fmt.Printf("Cleaning cache, Lenght:%d", len(cache.entries))
		cache.mux.Unlock()
	}
}
