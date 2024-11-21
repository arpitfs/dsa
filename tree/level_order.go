package tree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func TreeSetUp() *Node {
	input := []int{1, 2, 3, 4, 5, -1, 6, -1, -1, 7, -1, -1, -1, -1, -1}
	root := buildLevelOrderTree(input)
	traverseLevelOrder(root)
	return root

}

func buildLevelOrderTree(input []int) *Node {
	root := &Node{data: input[0]}
	queue := []*Node{root}

	pointerToInput := 1

	for len(queue) > 0 {
		currentNode := queue[0]
		if len(queue) > 0 {
			queue = queue[1:]
		}
		childFirst := input[pointerToInput]
		childSecond := input[pointerToInput+1]

		if childFirst != -1 {
			leftNode := &Node{data: childFirst}
			currentNode.left = leftNode
			queue = append(queue, leftNode)
		}

		if childSecond != -1 {
			rightNode := &Node{data: childSecond}
			currentNode.right = rightNode
			queue = append(queue, rightNode)
		}

		pointerToInput = pointerToInput + 2
	}

	return root
}

func traverseLevelOrder(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	queue = append(queue, nil)

	for len(queue) > 0 {
		currentNode := queue[0]
		if currentNode == nil {
			fmt.Println()
			queue = queue[1:]

			if len(queue) != 0 {
				queue = append(queue, nil)
			}
		} else {
			queue = queue[1:]
			fmt.Print(currentNode.data, " ")
			if currentNode.left != nil {
				queue = append(queue, currentNode.left)
			}

			if currentNode.right != nil {
				queue = append(queue, currentNode.right)
			}
		}
	}

}
