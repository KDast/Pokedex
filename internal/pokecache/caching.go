package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval int) cache {
	var newCache cache
	newCache.entry = map[string]cacheEntry{}
	newCache.mu = new(sync.Mutex)
	timer := time.Duration(interval) * time.Second
	go newCache.reapLoop(timer)
	return newCache
}

type cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

func (c cache) Add(key string, value []byte) {
	var cacheAdd cacheEntry
	cacheAdd = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.mu.Lock()
	c.entry[key] = cacheAdd
	c.mu.Unlock()

}
func (c cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock() //will unlock the mu when the function exit no matter how it exit
	k, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	value := k.val
	return value, true
}

func (c cache) reapLoop(timer time.Duration) {
	ticker := time.NewTicker(timer)
	for {
		<-ticker.C
		c.mu.Lock()
		for k := range c.entry {
			age := time.Since(c.entry[k].createdAt)
			if age >= timer {
				delete(c.entry, k)
			}
		}
		c.mu.Unlock()
	}

}
