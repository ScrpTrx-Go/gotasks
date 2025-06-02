package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type cacheItem struct {
	value     int
	createdAt time.Time
	ttl       time.Duration
}

type shard struct {
	mu    sync.RWMutex
	items map[int]cacheItem
}

// TTLCache — основной кэш
type TTLCache struct {
	shards     []*shard
	shardCount int
	startTime  time.Time
}

// NewTTLCache — конструктор
func NewTTLCache(shardCount int) *TTLCache {
	shards := make([]*shard, shardCount)
	for i := range shards {
		shards[i] = &shard{items: make(map[int]cacheItem)}
	}
	return &TTLCache{
		shards:     shards,
		shardCount: shardCount,
		startTime:  time.Now(),
	}
}

func (c *TTLCache) log(format string, a ...any) {
	elapsed := time.Since(c.startTime).Truncate(time.Millisecond)
	prefix := fmt.Sprintf("[%s] ", elapsed)
	fmt.Printf(prefix+format+"\n", a...)
}

func (c *TTLCache) getShard(key int) *shard {
	return c.shards[key%c.shardCount]
}

func (c *TTLCache) Put(key, value int, ttl time.Duration) {
	s := c.getShard(key)
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.items[key]; ok {
		c.log("[PUT]: key %d already exists, overwriting", key)
	} else {
		c.log("[PUT]: key %d added", key)
	}

	time.Sleep(1 * time.Second)
	s.items[key] = cacheItem{value: value, createdAt: time.Now(), ttl: ttl}
}

func (c *TTLCache) Get(key int) (int, bool) {
	s := c.getShard(key)
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[key]
	if !ok || time.Since(item.createdAt) > item.ttl {
		c.log("[GET]: key %d not found or expired", key)
		return 0, false
	}
	c.log("[GET]: key %d found", key)
	return item.value, true
}

func (c *TTLCache) Delete(key int) {
	s := c.getShard(key)
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.items[key]; ok {
		c.log("[DELETE]: key %d deleted", key)
		delete(s.items, key)
	} else {
		c.log("[DELETE]: key %d not found", key)
	}
}

func (c *TTLCache) Clean() {
	for _, s := range c.shards {
		s.mu.Lock()
		for key, item := range s.items {
			if time.Since(item.createdAt) > item.ttl {
				delete(s.items, key)
			}
		}
		s.mu.Unlock()
	}
	c.log("[CLEANER]: cache cleaned")
}

func (c *TTLCache) PrintStats() {
	for i, s := range c.shards {
		s.mu.RLock()
		fmt.Printf("Shard %d size: %d\n", i, len(s.items))
		s.mu.RUnlock()
	}
}

func main() {
	cache := NewTTLCache(4)

	var wg sync.WaitGroup

	// Очистка каждые 2 секунды
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			cache.Clean()
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Put(i, rand.Intn(100), 2*time.Second)
		}(i)
	}

	time.Sleep(1 * time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if val, ok := cache.Get(3); ok {
			fmt.Println("GET 3 =", val)
		} else {
			fmt.Println("GET 3 not found")
		}
	}()

	time.Sleep(3 * time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Delete(3)
	}()

	wg.Wait()
	cache.PrintStats()
}
