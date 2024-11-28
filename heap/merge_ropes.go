package heap

import (
	"container/heap"
	"fmt"
)

func getMinRopes() {
	input := []int{4, 3, 2, 6}

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, val := range input {
		heap.Push(minHeap, val)
	}

	minHeap.printRopes()

}

func (h *MinHeap) printRopes() {
	result := 0
	for h.Len() > 1 {
		data1 := heap.Pop(h)
		fmt.Println("data1: ", data1)
		data2 := heap.Pop(h)
		data3 := data1.(int) + data2.(int)
		result += data3
		heap.Push(h, data3)
	}

	fmt.Println("Result:", result)
}
