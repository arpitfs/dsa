package linkedlist

func RemoveEnd() {
	data := []int{1, 2, 3, 4, 5}
	n := 2
	head := createList(data)
	PrintList(head)
	reversed := removeNode(head, n)
	PrintList(reversed)
}

func removeNode(head *Node, n int) *Node {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.next
	}

	nodeToBeDeleted := length - n + 1 // 4
	if nodeToBeDeleted == 1 {
		return head.next
	}
	current = head
	for i := 1; i < length-n; i++ {
		current = current.next
	}
	current.next = current.next.next

	return head

}
