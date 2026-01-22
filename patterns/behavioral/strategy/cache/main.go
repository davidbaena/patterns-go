package main

func main() {
	lru := LRU{}
	cache := NewCache(lru)

	cache.add("A", "1")
	cache.add("B", "2")

	cache.add("C", "3") // Evict

	lfu := LFU{}
	cache.SetEvictAlg(lfu)
	cache.add("D", "4") // Evict

	fifo := Fifo{}
	cache.SetEvictAlg(fifo)
	cache.add("E", "5") // Evict
}
