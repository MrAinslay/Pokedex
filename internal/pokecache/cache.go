package pokecache

import (
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Time{},
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.cacheMap[key]
	if !ok {
		return []byte{}, ok
	}
	return val.val, ok
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.interval * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for key, cacheEntry := range c.cacheMap {
				creationTime := cacheEntry.createdAt
				t := time.Since(creationTime)
				if t >= c.interval*time.Second {
					c.mu.Lock()
					delete(c.cacheMap, key)
					c.mu.Unlock()
				}
			}
		}

	}
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		cacheMap: map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}
}
