package main

import (
	"arrayMain"
	"fmt"
	"ll"
	"lru"
	"stringMain"
)

func main() {
	fmt.Println("Start Running DSA Programs")
	arrayMain.Array()
	fmt.Println("Completed Running Array Programs")
	stringMain.String()
	fmt.Println("Completed Running String Programs")
	lru.LRUCache()
	fmt.Println("Printed LRU Cache")
	ll.DetectCycle()
}
