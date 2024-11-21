package tree

import (
	"fmt"
	"math"
)

func balanced() {
	root := TreeSetUp()

	_, isBalance := isBalanced(root)
	fmt.Println("Is Balanced: ", isBalance)
}

func isBalanced(root *Node) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, isLeftTree := isBalanced(root.left)
	rightHeight, isRightTree := isBalanced(root.right)
	h := maximum(leftHeight, rightHeight)

	if math.Abs(float64(leftHeight)-float64(rightHeight)) <= 1 && isLeftTree && isRightTree {
		return h, true
	}
	return h, false
}
