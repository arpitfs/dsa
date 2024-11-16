package linkedlist

import "fmt"

func mergeKSort() {
	record1 := []int{2, 3, 5}
	head1 := createList(record1)
	record2 := []int{1, 2, 6}
	head2 := createList(record2)
	record3 := []int{6, 8, 9}
	head3 := createList(record3)
	record4 := []int{4, 7, 9}
	head4 := createList(record4)

	list := []*Node{head1, head2, head3, head4}
	head := mergeSortedList(list)
	fmt.Println("merge K sorted")
	PrintList(head)
}

func mergeSortedList(list []*Node) *Node {
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

func mergeTwoSortedList(node1, node2 *Node) *Node {
	dummy := &Node{}
	current := dummy
	for node1 != nil && node2 != nil {
		if node1.value < node2.value {
			current.next = node1
			current = current.next
			node1 = node1.next
		} else {
			current.next = node2
			current = current.next
			node2 = node2.next
		}

		if node1 == nil {
			current.next = node2
		}

		if node2 == nil {
			current.next = node1
		}
	}
	return dummy.next
}
