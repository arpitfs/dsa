package heap

import (
	"container/heap"
	"fmt"
)

func runningMedians() {

	// work with float type
	input := []int{10, 5, 2, 3, 0, 12, 18, 20, 22}

	rightHeap := &MinHeap{}
	leftHeap := &MaxHeap{}
	heap.Init(leftHeap)
	heap.Init(rightHeap)
	median := input[0]
	heap.Push(leftHeap, input[0])
	fmt.Print("Medians: ", median, " ,")

	for i := 1; i < len(input); i++ {
		if leftHeap.Len() > rightHeap.Len() {
			if input[i] < median {
				data := (*leftHeap)[0]
				heap.Push(rightHeap, data)
				heap.Pop(leftHeap)
				heap.Push(leftHeap, input[i])
			} else {
				heap.Push(rightHeap, input[i])
			}
			leftData := heap.Pop(leftHeap)
			rightData := heap.Pop(rightHeap)
			median = (leftData.(int) + rightData.(int)) / 2
		} else if leftHeap.Len() == rightHeap.Len() {
			if input[0] < median {
				heap.Push(leftHeap, input[i])
				median = heap.Pop(leftHeap).(int)
			} else {
				heap.Push(rightHeap, input[i])
				median = heap.Pop(rightHeap).(int)
			}
		} else {
			if input[i] < median {
				heap.Push(leftHeap, input[i])

			} else {
				data := (*rightHeap)[0]
				heap.Push(leftHeap, data)
				heap.Pop(rightHeap)
				heap.Push(rightHeap, input[i])
			}
			leftData := heap.Pop(leftHeap)
			rightData := heap.Pop(rightHeap)
			median = (leftData.(int) + rightData.(int)) / 2
		}

		fmt.Print(median, " ,")
	}
}
