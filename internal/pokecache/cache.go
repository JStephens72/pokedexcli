package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	timeout time.Duration
	mu      sync.Mutex
}

// NewClient
func NewCache(cleanupInterval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		timeout: cleanupInterval,
	}

	go c.reapLoop(cleanupInterval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	//lazy expiration
	//return a negative cache hit if the entry has expired,
	//but has not yet been cleaned up by the reapLoop
	if time.Since(entry.createdAt) > c.timeout {
		delete(c.entries, key)
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) RemoveExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for key, entry := range c.entries {
		if now.Sub(entry.createdAt) > c.timeout {
			delete(c.entries, key)
		}
	}
}

func (c *Cache) reapLoop(cleanupInterval time.Duration) {
	ticker := time.NewTicker(cleanupInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.RemoveExpired()
	}
}
