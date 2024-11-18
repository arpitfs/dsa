package main

import (
	"array"
	"bitwise"
	"dp"
	"fmt"
	"linkedlist"
	"lru"
	binarysearch "main/binary-search"
	twopointer "main/two-pointer"
	"string"
)

func main() {
	fmt.Println("Start Running DSA Programs")
	array.Array()
	fmt.Println("Completed Running Array Programs")
	string.String()
	fmt.Println("Completed Running String Programs")
	lru.LRUCache()
	fmt.Println("Printed LRU Cache")
	linkedlist.ListMain()
	dp.DpMain()
	bitwise.Bitwise()
	twopointer.GetHouses()
	binarysearch.RotatedSearch()
}
