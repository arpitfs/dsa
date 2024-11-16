package helper

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func CreateList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	head := &Node{Value: values[0]}
	current := head
	for _, val := range values[1:] {
		newNode := &Node{Value: val}
		current.Next = newNode
		current = newNode
	}
	return head
}

func PrintList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}
