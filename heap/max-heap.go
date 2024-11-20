package heap

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxHeap() {
	h := &MaxHeap{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	heap.Init(h)
	fmt.Printf("Max-Heap: %v\n", *h)
	// Push a new element to the heap
	heap.Push(h, 10)

	// Pop the largest element
	fmt.Printf("Max-Heap: %v\n", *h)
	fmt.Printf("Popped: %d\n", heap.Pop(h))
	fmt.Printf("Max-Heap after pop: %v\n", *h)
}
