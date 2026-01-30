package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	CacheEntriesMap map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		CacheEntriesMap : make(map[string]cacheEntry),
		interval : interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CacheEntriesMap[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.CacheEntriesMap[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap () {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	
	for key, val := range c.CacheEntriesMap {
		if now.Sub(val.createdAt) > c.interval {
			delete(c.CacheEntriesMap, key)
		}
	}
}
