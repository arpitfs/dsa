package tree

import "fmt"

type TNode struct {
	value int
	left  *TNode
	right *TNode
}

func treeOperations() {
	var root *TNode
	values := []int{50, 30, 20, 40, 70, 60, 80}
	for _, value := range values {
		root = InsertBST(root, value)
	}

	inOrderBSTravers(root)
	fmt.Println()
	deleteNode(root, 20)
	inOrderBSTravers(root)
	fmt.Println()
}

func deleteNode(root *TNode, value int) *TNode {
	if root == nil {
		return nil
	}

	if value < root.value {
		root.left = deleteNode(root.left, value)
	} else if value > root.value {
		root.right = deleteNode(root.right, value)
	} else {
		if root.left == nil && root.right == nil {
			return nil
		}

		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			minNode := findMinValue(root.right)
			root.value = minNode.value
			root.right = deleteNode(root.right, minNode.value)
		}

	}

	return root

}

func findMinValue(root *TNode) *TNode {
	c := root
	for c.left != nil {
		c = c.left
	}
	return c
}
func inOrderBSTravers(root *TNode) {
	if root != nil {
		inOrderBSTravers(root.left)
		fmt.Printf("%d ", root.value)
		inOrderBSTravers(root.right)
	}

}

func InsertBST(root *TNode, value int) *TNode {
	if root == nil {
		root.value = value
	}

	if value < root.value {
		root.left = InsertBST(root, value)
	} else {
		root.right = InsertBST(root, value)
	}

	return root
}
