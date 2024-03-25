package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu      *sync.RWMutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries: make(map[string]CacheEntry),
		mu:      &sync.RWMutex{},
	}
	go cache.readLoop(interval)
	return &newCache
}

func (cache *Cache) Add(key string, val []byte) {
	var entry CacheEntry = CacheEntry{createdAt: time.Now(),
		val: val,
	}
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	entry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache) readLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		cache.mu.RLock()
		for key, entry := range cache.entries {
			if time.Since(entry.createdAt) > interval {
				cache.mu.Lock()
				delete(cache.entries, key)
				cache.mu.Unlock()
			}
		}
		cache.mu.RUnlock()
	}
}
