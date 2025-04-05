package blind75

func mergeKSortedList() {
	list1 := []int{2, 3, 5}
	head1 := CreateList(list1)
	list2 := []int{1, 2, 6}
	head2 := CreateList(list2)
	list3 := []int{6, 8, 9}
	head3 := CreateList(list3)

	list := []*Node{head1, head2, head3}

	head := mergeSortedKList(list)
	PrintList(head)
}

func mergeSortedKList(list []*Node) *Node {
	if len(list) == 0 {
		return nil
	}

	interval := 1
	for interval < len(list) {
		for i := 0; i+interval < len(list); i = i + interval*2 {
			list[i] = mergeKTwoList(list[i], list[i+interval])
		}

		interval = interval * 2
	}
	return list[0]
}

func mergeKTwoList(list1, list2 *Node) *Node {
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
