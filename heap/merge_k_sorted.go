package heap

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value, arrayIndex, elementIndex int
}

type MinHeapSorted []Item

func (h MinHeapSorted) Len() int           { return len(h) }
func (h MinHeapSorted) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeapSorted) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapSorted) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeapSorted) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeSorted() {
	input := [][]int{{10, 15, 20, 30}, {2, 5, 8, 14, 24}, {0, 11, 60, 90}}
	mergedArray := MergeKSortedArrays(input)
	fmt.Println("Merged Array:", mergedArray)
}

func MergeKSortedArrays(arrays [][]int) []int {
	result := []int{}
	minHeap := &MinHeapSorted{}

	for i := 0; i < len(arrays); i++ {
		if len(arrays[i]) > 0 {
			heap.Push(minHeap, Item{value: arrays[i][0], arrayIndex: i, elementIndex: 0})
		}
	}

	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(Item)
		result = append(result, item.value)

		if item.elementIndex+1 < len(arrays[item.arrayIndex]) {
			nextElement := arrays[item.arrayIndex][item.elementIndex+1]
			heap.Push(minHeap, Item{value: nextElement, arrayIndex: item.arrayIndex, elementIndex: item.elementIndex + 1})
		}
	}

	return result
}
