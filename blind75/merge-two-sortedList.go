package blind75

func mergeSortedList() {
	list1 := []int{2, 3, 5}
	head1 := CreateList(list1)
	list2 := []int{1, 2, 6}
	head2 := CreateList(list2)

	head := mergeTwoList(head1, head2)
	PrintList(head)
}

func mergeTwoList(list1, list2 *Node) *Node {
	dummyNode := &Node{}
	current := dummyNode

	for list1 != nil && list2 != nil {
		if list1.data < list2.data {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		} else {
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		}

		if list1 == nil {
			current.Next = list2
		}

		if list2 == nil {
			current.Next = list1
		}
	}
	return dummyNode.Next
}
