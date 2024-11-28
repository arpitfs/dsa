package linkedlist

import "main/linkedlist/helper"

func RemoveEnd() {
	data := []int{1, 2, 3, 4, 5}
	n := 2
	head := helper.CreateList(data)
	helper.PrintList(head)
	reversed := removeNode(head, n)
	helper.PrintList(reversed)
}

func removeNode(head *helper.Node, n int) *helper.Node {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	nodeToBeDeleted := length - n + 1 // 4
	if nodeToBeDeleted == 1 {
		return head.Next
	}
	current = head
	for i := 1; i < length-n; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next

	return head

}
