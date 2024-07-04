package config

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)

	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for k, v := range c.cache {
		if now.Sub(v.createdAt) > interval {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) Delete(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	delete(c.cache, key)
}

func (c *Cache) Reset() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache = make(map[string]cacheEntry)
}
