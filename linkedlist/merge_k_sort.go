package linkedlist

import (
	"fmt"
	"linkedlist/helper"
)

func mergeKSort() {
	record1 := []int{2, 3, 5}
	head1 := helper.CreateList(record1)
	record2 := []int{1, 2, 6}
	head2 := helper.CreateList(record2)
	record3 := []int{6, 8, 9}
	head3 := helper.CreateList(record3)
	record4 := []int{4, 7, 9}
	head4 := helper.CreateList(record4)

	list := []*helper.Node{head1, head2, head3, head4}
	head := mergeSortedList(list)
	fmt.Println("merge K sorted")
	helper.PrintList(head)
}

func mergeSortedList(list []*helper.Node) *helper.Node {
	if len(list) == 0 {
		return nil
	}

	interval := 1
	for interval < len(list) {
		for i := 0; i+interval < len(list); i = i + interval*2 {
			list[i] = mergeTwoSortedList(list[i], list[i+interval])
		}
		interval = interval * 2
	}
	return list[0]

}

func mergeTwoSortedList(node1, node2 *helper.Node) *helper.Node {
	dummy := &helper.Node{}
	current := dummy
	for node1 != nil && node2 != nil {
		if node1.Value < node2.Value {
			current.Next = node1
			current = current.Next
			node1 = node1.Next
		} else {
			current.Next = node2
			current = current.Next
			node2 = node2.Next
		}

		if node1 == nil {
			current.Next = node2
		}

		if node2 == nil {
			current.Next = node1
		}
	}
	return dummy.Next
}
