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
	for _, val := range values[1:] {
		newNode := &Node{value: val}
		current.next = newNode
		current = newNode
	}
	return head
}

func PrintList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}
