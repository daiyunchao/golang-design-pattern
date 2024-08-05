package evict

import (
	"testing"
)

func TestCache(t *testing.T) {
	//fifo 算法
	fifoAlgo := &FIFOEvictionAlgo{}
	cache := InitCache()
	cache.SetEvictAlgo(fifoAlgo)
	cache.AddCache("one", "1")
	cache.AddCache("two", "2")
	cache.AddCache("three", "3")
	cache.AddCache("four", "4")
}
