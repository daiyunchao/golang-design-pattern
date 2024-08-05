package evict

type Cache struct {
	storage     map[string]string
	evictAlgo   EvictionAlgo
	capacity    int
	maxCapacity int
}

func InitCache() *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:     storage,
		evictAlgo:   nil,
		capacity:    0,
		maxCapacity: 2,
	}
}

func (c *Cache) SetEvictAlgo(evictAlgo EvictionAlgo) {
	c.evictAlgo = evictAlgo
}

func (c *Cache) AddCache(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) GetCache(key string) (string, bool) {
	val, ok := c.storage[key]
	return val, ok
}

func (c *Cache) evict() {
	c.evictAlgo.evict(c)
	c.capacity--
}
