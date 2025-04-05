package blind75

func removeEnd() {
	input := []int{1, 2, 3, 4, 5}
	n := 2
	head := CreateList(input)
	removedHead := removeNodeEnd(head, n)
	PrintList(removedHead)
}

func removeNodeEnd(head *Node, n int) *Node {
	if head == nil {
		return nil
	}

	l := 0
	c := head
	for c != nil {
		l++
		c = c.Next
	}

	nodeToBeDeleted := l - n + 1
	if nodeToBeDeleted == 1 {
		return head.Next
	}

	for i := 1; i < l-n; i++ {
		c = c.Next
	}
	c.Next = c.Next.Next
	return head
}
