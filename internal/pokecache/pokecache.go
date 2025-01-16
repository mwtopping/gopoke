package pokecache

import (
	"sync"
	"time"

	"fmt"
)

type CacheEntry struct {
	createdAt time.Time
	Val       []byte
}

type Cache struct {
	data map[string]CacheEntry
	mu   sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	// done := make(chan bool)

	cache := Cache{data: make(map[string]CacheEntry)}

	go func() {
		for {
			<-ticker.C
			fmt.Println("Tick")
			cache.ReadLoop(interval)
		}
	}()

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	newEntry := CacheEntry{createdAt: time.Now(), Val: val}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = newEntry

	return
}

func (c *Cache) Get(key string) (CacheEntry, bool) {

	if val, ok := c.data[key]; ok {
		return val, true
	}

	return CacheEntry{}, false
}

func (c *Cache) ReadLoop(interval time.Duration) {
	for key := range c.data {
		if c.data[key].createdAt.Add(interval).Before(time.Now()) == true {
			// entry is old
			fmt.Println("Found old key", key, "deleting")
			delete(c.data, key)
		}
	}
}
