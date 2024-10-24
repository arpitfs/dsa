package lru

import "fmt"

func LRUCache() {
	cache := CreateLRUCache(2)
	cache.addValueFromCache(1, 1)
	cache.addValueFromCache(2, 2)
	fmt.Println(cache.getValueFromCache(1)) // returns 1
	cache.addValueFromCache(3, 3)           // evicts key 2
	fmt.Println(cache.getValueFromCache(2)) // returns 0 (not found)
	cache.addValueFromCache(4, 4)           // evicts key 1
	fmt.Println(cache.getValueFromCache(1)) // returns 0 (not found)
	fmt.Println(cache.getValueFromCache(3)) // returns 3
	fmt.Println(cache.getValueFromCache(4)) // retuurn 4
}
