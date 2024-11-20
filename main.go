package main

import (
	"array"
	"bitwise"
	"dp"
	"linkedlist"
	"lru"
	binarysearch "main/binary-search"
	"main/hashing"
	"main/heap"
	"main/recursion"
	twopointer "main/two-pointer"
	"string"
)

func main() {
	array.ArrayProblems()
	string.StringProblems()
	lru.LRUCache()
	linkedlist.ListProblems()
	dp.DpProblems()
	bitwise.BitwiseProblems()
	twopointer.TwoPointerProblems()
	binarysearch.BinarySearchProblems()
	recursion.RecursionProblems()
	hashing.HashingProblem()
	heap.HeapProblems()
}
