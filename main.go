package main

import (
	"main/array"
	binarysearch "main/binary-search"
	"main/bitwise"
	"main/dp"
	"main/hashing"
	"main/heap"
	"main/linkedlist"
	"main/lru"
	"main/recursion"
	"main/string"
	"main/tree"
	twopointer "main/two-pointer"
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
	tree.TreeProblems()
}
