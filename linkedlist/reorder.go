package linkedlist

func reorder() {
	data := []int{1, 2, 3, 4, 5}
	head := createList(data)
	reorderListHead := reorderList(head)
	PrintList(reorderListHead)
}

func reorderList(head *Node) *Node {
	if head.next == nil || head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	head1 := head
	head2 := slow.next
	slow.next = nil
	head2 = reverseList(head2)
	current := head

	for head1 != nil {
		temp := head1.next
		current.next = head1
		head1.next = head2
		current = head2
		head1 = temp
		if head2 != nil {
			head2 = head2.next
		}
	}
	return head
}

func reverseList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}

	return prev
}
