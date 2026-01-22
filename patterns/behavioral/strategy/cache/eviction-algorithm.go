package main

import "fmt"

type EvictionAlgorithm interface {
	evict(c *Cache)
}

type LFU struct{}

func (LFU) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}
