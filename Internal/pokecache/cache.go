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

func NewCache() {

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheEntry[key] = cacheEntry{val = val, createdAt = time.now()}
	c.mu.Unlock()
}
	
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, exists := c.cacheEntry[key]
	if exists {
		return c.cacheEntry[key].val, true
	} else {
		return []byte{}, false
	}
	
}

func (c *Cache) reapLoop() {

}
