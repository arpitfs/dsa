package tree

import "fmt"

func printKLevel() {
	root := TreeSetUp()
	k := 2
	PrintKLevelTree(root, k)
}

func PrintKLevelTree(root *Node, k int) {
	if root == nil {
		return
	}
	// 1, 2, 3, 4, 5, -1, 6, -1, -1, 7, -1, -1, -1, -1, -
	queue := []*Node{root}
	level := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		levelNodes := []int{}

		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			if level == k {
				levelNodes = append(levelNodes, currentNode.data)
			}

			if currentNode.left != nil {
				queue = append(queue, currentNode.left)
			}

			if currentNode.right != nil {
				queue = append(queue, currentNode.right)
			}
		}

		if level == k {
			fmt.Println("Nodes at level", k, ":", levelNodes)
			return
		}
		level++

	}

}
