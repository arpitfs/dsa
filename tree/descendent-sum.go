package tree

import "fmt"

func sum() {
	root := TreeSetUp()
	getSum(root)
	fmt.Println("Sum Descendent")
	traverseLevelOrder(root)

}

func getSum(root *Node) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return root.data
	}

	leftSum := getSum(root.left)
	rightSum := getSum(root.right)

	temp := root.data

	root.data = leftSum + rightSum
	return temp + root.data

}
