package ll

import "fmt"

func ReverseList() {
	data := []int{2, 4, 5, 1}
	head := createList(data)
	PrintList(head)
	fmt.Println()
	prev := reverse(head)
	PrintList(prev)
}

func reverse(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

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
