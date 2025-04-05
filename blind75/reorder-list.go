package blind75

func reorderListMain() {
	input := []int{1, 2, 3, 4, 5}
	head := CreateList(input)
	reorderedList := reorderList(head)
	PrintList(reorderedList)
}

func reorderList(head *Node) *Node {
	if head.Next == nil || head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head1 := head
	head2 := slow.Next
	slow.Next = nil
	head2 = reverseNodes(head2)
	current := head

	for head1 != nil {
		temp := head1.Next
		current.Next = head1
		head1.Next = head2
		current = head2
		head1 = temp
		if head2 != nil {
			head2 = head2.Next
		}
	}
	return head
}

func reverseNodes(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}

	return prev
}
