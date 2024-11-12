package ll

import "fmt"

type Node struct {
	value int
	next  *Node
}

func createList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	head := &Node{value: values[0]}
	current := head
	for _, val := range values {
		newNode := &Node{value: val}
		current.next = newNode
		current = newNode
	}
	return head
}

func PrintList(head *Node) {
	for head.next != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

func DetectCycle() {
	data := []int{1, 2, 34, 5, 1}
	head := createList(data)
	PrintList(head)
}
