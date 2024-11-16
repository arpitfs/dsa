package linkedlist

import (
	"fmt"
	"linkedlist/helper"
)

func ReverseList() {
	data := []int{2, 4, 5, 1}
	head := helper.CreateList(data)
	helper.PrintList(head)
	fmt.Println()
	prev := reverse(head)
	helper.PrintList(prev)
}

func reverse(head *helper.Node) *helper.Node {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *helper.Node
	current := head

	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	return prev
}
