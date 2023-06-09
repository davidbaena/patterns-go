package main

import "fmt"

type LRU struct{}

func (LRU) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}
