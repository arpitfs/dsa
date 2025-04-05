package blind75

func revereList() {
	input := []int{1, 2, 3, 4, 5}
	head := CreateList(input)
	PrintList(head)
	reversedNode := reverse(head)
	PrintList(reversedNode)
}

func reverse(head *Node) *Node {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

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
