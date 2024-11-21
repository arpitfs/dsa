package tree

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func Insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{value: value}
	}

	if value < root.value {
		root.left = Insert(root.left, value)
	} else {
		root.right = Insert(root.right, value)
	}

	return root

}

func InOrderTraverse(root *TreeNode) {
	if root != nil {
		InOrderTraverse(root.left)
		fmt.Println(root.value, " ")
		InOrderTraverse(root.right)
	}
}

func PreOrderTraverse(root *TreeNode) {
	if root != nil {
		fmt.Println(root.value, " ")
		PreOrderTraverse(root.left)
		PreOrderTraverse(root.right)
	}
}

func PostOrderTraverse(root *TreeNode) {
	if root != nil {
		PostOrderTraverse(root.left)
		PostOrderTraverse(root.right)
		fmt.Println(root.value, " ")
	}
}

func Tree() {
	root := &TreeNode{value: 10}
	Insert(root, 5)
	Insert(root, 15)
	Insert(root, 3)
	Insert(root, 7)
	Insert(root, 12)
	Insert(root, 18)

	fmt.Println("In-order traversal:")
	InOrderTraverse(root)
	fmt.Println()

	fmt.Println("Pre-order traversal:")
	PreOrderTraverse(root)
	fmt.Println()

	fmt.Println("Post-order traversal:")
	PostOrderTraverse(root)
	fmt.Println()
}
