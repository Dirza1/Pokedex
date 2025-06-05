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
	_, exisits := c.cacheEntry[key]
	if exisits{
		c.cacheEntry[key].val = val
		c.cacheEntry[key].createdAt = time.Now()
		c.mu.Unlock()
	} else {
		c.cacheEntry[key] := make(cacheEntry)
		c.cacheEntry[key].val = val
		c.cacheEntry[key].createdAt = time.Now()
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
