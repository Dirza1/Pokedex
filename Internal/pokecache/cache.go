package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{}
	c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntry[key] = cacheEntry{val: val,
		createdAt: time.Now()}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cacheEntry[key]
	if exists {
		return entry.val, true
	} else {
		return []byte{}, false
	}

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		currentTime := time.Now()
		for key, entry := range c.cacheEntry {
			duration := currentTime.Sub(entry.createdAt)
			if duration > interval {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
