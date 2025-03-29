package practice

import "fmt"

type BST struct {
	value int
	left  *BST
	right *BST
}

func AddValue(root *BST, value int) *BST {
	if root == nil {
		return &BST{value: value}
	}

	if value < root.value {
		root.left = AddValue(root.left, value)
	} else {
		root.right = AddValue(root.right, value)
	}

	return root
}

func inOrderTraversal(root *BST) {
	if root != nil {
		inOrderTraversal(root.left)
		fmt.Println(root.value)
		inOrderTraversal(root.right)
	}
}

func deleteNode(root *BST, value int) *BST {
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
			return root.right
		} else {
			min := findMin(root.right)
			root.value = min.value
			root.left = deleteNode(root.right, min.value)
		}
	}
	return root
}

func findMin(root *BST) *BST {
	c := root
	if c.left != nil {
		c = c.left
	}
	return c
}

func binarySearchTree() {
	root := &BST{value: 2}
	fmt.Println(root)
	AddValue(root, 3)
	AddValue(root, 1)
	AddValue(root, 4)
	inOrderTraversal(root)
	fmt.Println()
	deleteRoot := deleteNode(root, 3)
	inOrderTraversal(deleteRoot)

}
