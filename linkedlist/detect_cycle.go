package linkedlist

import (
	"fmt"
	"linkedlist/helper"
)

func DetectCycle() {
	data := []int{1, 2, 34, 5, 1}
	head := helper.CreateList(data)
	helper.PrintList(head)
	fmt.Println(hasCycle(head))
}

func hasCycle(head *helper.Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	result := false
	for fast != nil && fast.Next != nil {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			result = true
			break
		}

	}
	return result
}