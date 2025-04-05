package blind75

import "fmt"

type Node struct {
	data int
	Next *Node
}

func CreateList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	head := &Node{data: values[0]}
	current := head
	for _, v := range values {
		newNode := &Node{data: v}
		current.Next = newNode
		current = newNode
	}
	return head
}

func PrintList(head *Node) {
	current := head

	for current != nil {
		fmt.Println(current.data, " ->")
		current = current.Next
	}
	fmt.Println("nil")
}

func detetCycle() {
	input := []int{1, 2, 34, 5, 1}
	head := CreateList(input)
	PrintList(head)
	fmt.Println(checkCycle(head))
}

func checkCycle(head *Node) bool {
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
