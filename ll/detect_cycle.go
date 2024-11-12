package ll

func DetectCycle() {
	data := []int{1, 2, 34, 5, 1}
	head := createList(data)
	PrintList(head)
}

func hasCycle(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}
	slow := head
	fast := head
	result := false
	for fast != nil && fast.next != nil {
		if fast == nil || fast.next == nil {
			break
		}
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			result = true
			break
		}

	}
	return result
}
