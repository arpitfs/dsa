package heap

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minHeap() {
	h := &MinHeap{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	heap.Init(h)
	heap.Push(h, 10)

	fmt.Printf("Min-Heap: %v\n", *h)
	fmt.Printf("Popped: %d\n", heap.Pop(h))
	fmt.Printf("Min-Heap after pop: %v\n", *h)
}
