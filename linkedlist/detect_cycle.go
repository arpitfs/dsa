package linkedlist

import (
	"fmt"
	"main/linkedlist/helper"
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
	for fast != nil && fast.Next != nil {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}

	}
	return false
}
