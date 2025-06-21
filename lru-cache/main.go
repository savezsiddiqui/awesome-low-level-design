package main

import (
	"fmt"

	"example.com/lru-cache/lru"
)

func main() {
	lruCache := lru.NewLruCache(3)
	fmt.Println(lruCache.Get(5))
	lruCache.Put(5, 5)
	lruCache.Put(6, 5)
	lruCache.Put(7, 5)
	fmt.Println(lruCache.Get(7))
}
