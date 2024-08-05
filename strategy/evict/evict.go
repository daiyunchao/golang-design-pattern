package evict

import "github.com/marmotedu/iam/pkg/log"

// EvictionAlgo 驱逐算法
type EvictionAlgo interface {
	evict(cache *Cache)
}

type FIFOEvictionAlgo struct {
}

func (algo FIFOEvictionAlgo) evict(cache *Cache) {
	log.Infof("FIFO eviction algorithm")
}

type LRUEvictionAlgo struct{}

func (algo LRUEvictionAlgo) evict(cache *Cache) {
	log.Infof("LRU eviction algorithm")
}

type LFUEvictionAlgo struct{}

func (algo LFUEvictionAlgo) evict(cache *Cache) {
	log.Infof("LFU eviction algorithm")
}
