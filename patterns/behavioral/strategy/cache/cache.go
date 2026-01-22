package main

type Cache struct {
	storage     map[string]string
	evictAlg    EvictionAlgorithm
	capacity    int
	maxCapacity int
}

func NewCache(alg EvictionAlgorithm) Cache {
	storage := make(map[string]string)
	return Cache{
		storage:     storage,
		evictAlg:    alg,
		capacity:    0,
		maxCapacity: 2,
	}
}

func (c *Cache) SetEvictAlg(evictAlg EvictionAlgorithm) {
	c.evictAlg = evictAlg
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictAlg.evict(c)
	c.capacity--
}
